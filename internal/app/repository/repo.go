package repository

import (
	"context"
	"strconv"

	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/dsn"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func (r *Repository) RewriteDBNews(ctx context.Context, DBitems model.DBNews) error {
	var err error

	year, month, d := DBitems.PublishedAt.Date()

	sqlDate := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(d)

	err = r.db.Where("published_at = ?", sqlDate).Delete(&DBitems).Error
	if err != nil {
		log.WithError(err).Error("cant delete news")

	}

	err = r.db.Create(&DBitems).Error
	if err != nil {
		log.WithError(err).Error("cant create item")
	}

	return nil
}
