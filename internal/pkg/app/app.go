package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/config"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/model"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/service"
)

type IService interface {
	ParseShortNews(ctx context.Context, limit int, offset int) ([]model.NewsItems, error)
	ParseFullNews(ctx context.Context, slug string) (model.FullNewsItem, error)
	JoinNewsInfo(ctx context.Context, shortNewsItem model.NewsItems, fullNewsItem model.FullNewsItem) model.News
	WriteDBNews(ctx context.Context, news model.News) error
	ReadDBNews(ctx context.Context, date time.Time) []model.News
	GetTotal(ctx context.Context) int
}

type App struct {
	ctx     context.Context
	service IService
}

func New(ctx context.Context) (*App, error) {
	app := &App{
		ctx: ctx,
	}

	srv, err := service.New(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Can`t create service")
	}

	app.service = srv
	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	newsCfg := config.FromContext(ctx).BMSTUNewsConfig

	for {
		shortNewsItems, err := a.service.ParseShortNews(ctx, 30, 0)
		if err != nil {
			log.WithError(err).Error("can`t parse news")

			continue
		}

		var fullNewsItem model.FullNewsItem

		for _, shortItem := range shortNewsItems {
			fullNewsItem, err = a.service.ParseFullNews(ctx, shortItem.Slug)
			if err != nil {
				log.WithError(err).Error("can`t parse full news")

				continue
			}

			newsItem := a.service.JoinNewsInfo(ctx, shortItem, fullNewsItem)

			err = a.service.WriteDBNews(ctx, newsItem)
		}

		time.Sleep(newsCfg.CronTimeout)
	}

	return nil
}

func (a *App) FirstParse(ctx context.Context) error {
	//реализовать парсинг всех старых новостей
	total := a.service.GetTotal(ctx)
	log.Info(total)
	limit := make(chan int, 1000)
	var wg sync.WaitGroup
	var b int

	for i := 0; i < total/100; i++ {

		for j := 1; j <= 5; j++ {
			b += 20
			limit <- b
		}

		for w := 1; w <= 1; w++ {
			wg.Add(1)
			go a.worker(ctx, &wg, limit)
		}

		wg.Wait()
	}

	close(limit)
	return nil
}

func (a *App) worker(ctx context.Context, wg *sync.WaitGroup, limit <-chan int) {
	defer wg.Done()
	offset := <-limit - 20
	fmt.Printf("limit - %v, offset - %v \n", offset+20, offset)
	shortNewsItems, err := a.service.ParseShortNews(ctx, offset+20, offset)
	if err != nil {
		log.WithError(err).Error("can`t parse news")
		return
	}
	var fullNewsItem model.FullNewsItem

	for _, shortItem := range shortNewsItems {
		fullNewsItem, err = a.service.ParseFullNews(ctx, shortItem.Slug)
		if err != nil {
			log.WithError(err).Error("can`t parse full news")

			continue
		}

		newsItem := a.service.JoinNewsInfo(ctx, shortItem, fullNewsItem)

		err = a.service.WriteDBNews(ctx, newsItem)
	}
}
