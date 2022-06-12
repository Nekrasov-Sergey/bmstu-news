package app

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"go.uber.org/atomic"

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
		shortNewsItems, err := a.service.ParseShortNews(ctx, newsCfg.DayLimit, 0)
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
	cfg := config.FromContext(ctx).FirstParse
	workerCount := int(cfg.WorkerCount)
	var failedFullParseCount atomic.Int32
	var failedShortParseCount int
	// get total
	total := a.service.GetTotal(ctx)
	shortNews := make([]model.NewsItems, total+100, total+100)

	for i := 0; i < total/cfg.StepParseSlug+1; i++ {
		offset := i * cfg.StepParseSlug

		shortNewsItems, err := a.service.ParseShortNews(ctx, cfg.StepParseSlug, offset)
		if err != nil {
			log.WithError(err).Error("can`t parse news")
			failedShortParseCount++

			continue
		}

		index := offset
		for _, item := range shortNewsItems {
			shortNews[index] = item
			index++
		}
	}

	log.WithField("total", total).Info("successfully get total")

	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		oneWorkerNews := cap(shortNews) / workerCount
		start := i * oneWorkerNews
		end := (i + 1) * oneWorkerNews
		if end > cap(shortNews) {
			end = cap(shortNews)
		}
		arr := shortNews[start:end]

		go func(failCounter atomic.Int32, shortNews []model.NewsItems) {
			for _, shNews := range shortNews {
				fullNewsItem, err := a.service.ParseFullNews(ctx, shNews.Slug)
				if err != nil {
					log.WithError(err).Error("can`t parse full news")

					failCounter.Add(1)
					continue
				}

				newsItem := a.service.JoinNewsInfo(ctx, shNews, fullNewsItem)

				err = a.service.WriteDBNews(ctx, newsItem)

				time.Sleep(1 * time.Second)
			}
			wg.Done()
		}(failedFullParseCount, arr)
	}

	wg.Wait()
	log.WithField("full_parse_count", failedShortParseCount).
		WithField("full_parse_count", failedFullParseCount.Load()).
		Info("count of failed parse")

	return nil
}
