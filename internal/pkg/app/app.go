package app

import (
	"context"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/config"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/model"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/service"
	log "github.com/sirupsen/logrus"
	"time"
)

type IService interface {
	ParseNews(ctx context.Context, limit string, offset string) ([]model.NewsItems, error)
	ParseFullNews(ctx context.Context, slug string) (model.FullNewsItems, error)
	WriteDBNews(ctx context.Context, NewsItems model.NewsItems, FullNewsItems model.FullNewsItems) error
	ReadDBNews(ctx context.Context, date time.Time) []model.DBNews
}

type App struct {
	ctx     context.Context
	service IService
}

func New(ctx context.Context) (*App, error) {
	/*
		db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
		if err != nil {
			log.WithError(err).Println("Can`t open postgres connection")
			return err
		}
	*/

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
		NewsItems, err := a.service.ParseNews(ctx, "5", "0")
		if err != nil {
			log.WithError(err).Error("Can`t parse news")
		}

		var FullNewsItems model.FullNewsItems

		for _, news := range NewsItems {
			FullNewsItems, err = a.service.ParseFullNews(ctx, news.Slug)
			if err != nil {
				log.WithError(err).Error("Can`t parse full news")
			}
			err = a.service.WriteDBNews(ctx, news, FullNewsItems)
		}

		time.Sleep(newsCfg.CronTimeout)
	}
	/*c := news.New(ctx)
	c.GetNews()*/

	return nil
}
