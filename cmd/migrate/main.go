package main

import (
	"context"
	"os"
	"strings"

	_ "github.com/lib/pq"

	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/dsn"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/pressly/goose"
	log "github.com/sirupsen/logrus"
)

const (
	migrationsPath = "migrations"
	driver         = "postgres"
)

func main() {

	ctx := context.Background()
	log.Info("Starting migrations")

	// Читает переменные окружения
	err := godotenv.Load()
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("No .env file loaded")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = strings.ReplaceAll(os.Getenv("CI_PROJECT_NAME"), "-", "_")
	}
	if dbName == "" {
		dbName = "postgres"
	}
	dbDSN := dsn.FromEnv()

	// подключаемся к БД
	db, err := connect(dbDSN)
	if err != nil {
		log.WithContext(ctx).WithError(err).Fatalln(err)
	}
	log.Info("The database connection was established successfully")

	// устанавливаем свой логер
	goose.SetLogger(&gooseLogger{ctx: ctx})
	_ = goose.SetDialect(driver)

	// запускаем миграции
	log.Info("Upping migrations")
	err = goose.Up(db.DB(), migrationsPath)
	if err != nil {
		log.Errorf("Failed to migrate: %v", err)
	}
	/*
		photo_report       text[],
			similar_news_slug  text[],
			tags_title         text[]*/
	log.Info("DB migration completed")
}

// Выполняет подключение к БД
func connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Реализация интерфйса goose.Logger
type gooseLogger struct {
	ctx context.Context
}

func (gl *gooseLogger) Fatal(v ...interface{}) {
	log.Fatal(v...)
}
func (gl *gooseLogger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
func (gl *gooseLogger) Print(v ...interface{}) {
	log.Info(v...)
}
func (gl *gooseLogger) Println(v ...interface{}) {
	log.Info(v...)
}
func (gl *gooseLogger) Printf(format string, v ...interface{}) {
	log.Infof(format, v...)
}
