package repository

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err := r.db.Create(&newsItem).Error
		if err != nil {
			return err
		}
	}

	// todo: проверять идентично ли и перезаписывать только если есть различия или делать UPDATE
	err = r.db.Where("slug = ?", newsItem.Slug).Delete(&selectedNewsItem).Error
	if err != nil {
		return err
	}

	err = r.db.Create(&newsItem).Error
	if err != nil {
		return err
	}

	return nil
}
