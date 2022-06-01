package main

import (
	"context"
	"fmt"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/config"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/pkg/app"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("cant init config")

		os.Exit(2)
	}

	ctx = config.WrapContext(ctx, cfg)

	fmt.Println(cfg)
	application, err := app.New(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can`t create application")

		os.Exit(2)
	}

	err = application.FirstParse(ctx)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("can`t run application")

		os.Exit(2)
	}

}
