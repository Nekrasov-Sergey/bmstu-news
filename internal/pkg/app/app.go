package app

import (
	"context"
)

type App struct {
	ctx context.Context
}

func New(ctx context.Context) (*App, error) {
	return &App{}, nil
}

func (a *App) Run() error {
	return nil
}
