package repository

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"

	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/dsn"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/model"
)

type Repository struct {
	// корневой контекст
	ctx context.Context

	db *gorm.DB
}

func New(ctx context.Context) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn.FromEnv()), &gorm.Config{})
	if err != nil {
		log.WithError(err).Println("Cant open postgers connection")

		return nil, err
	}

	return &Repository{
		ctx: ctx,
		db:  db,
	}, nil
}

func (r *Repository) RewriteDBNews(ctx context.Context, newsItem model.News) error {
	var err error

	selectedNewsItem := model.News{}

	err = r.db.
		Where("slug = ?", newsItem.Slug).
		First(&selectedNewsItem).
		Limit(1).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err := r.db.Create(&newsItem).Error
		if err != nil {
			return err
		}
	}

	if reflect.DeepEqual(selectedNewsItem, newsItem) {
		log.Info("news completely equal")
	} else {
		err = r.db.Where("slug = ?", newsItem.Slug).Updates(&newsItem).Error //не знаю, работает или нет :/
		if err != nil {
			return err
		}
	}

	return nil
}
