package app

import (
	"context"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/pkg/clients/news"
)

type App struct {
	ctx context.Context
}

func New(ctx context.Context) (*App, error) {
	return &App{}, nil
}

func (a *App) Run(ctx context.Context) error {
	c := news.New(ctx)
	c.GetNews()

	return nil
}
