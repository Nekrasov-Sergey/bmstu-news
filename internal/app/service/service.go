package service

import (
	"context"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/model"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/repository"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/pkg/clients/news"
)

var monthFormat = map[string]time.Month{
	"января":   time.January,
	"февраля":  time.February,
	"марта":    time.March,
	"апреля":   time.April,
	"мая":      time.May,
	"июня":     time.June,
	"июля":     time.July,
	"августа":  time.August,
	"сентября": time.September,
	"октября":  time.October,
	"ноября":   time.November,
	"декабря":  time.December,
}

type Service struct {
	newsClient *news.Client

	repo *repository.Repository
}

func New(ctx context.Context) (*Service, error) {
	newsClient := news.New(ctx)
	repo, err := repository.New(ctx)
	if err != nil {
		return nil, err
	}
	return &Service{
		newsClient: newsClient,
		repo:       repo,
	}, nil
}

func (s *Service) JoinNewsInfo(ctx context.Context,
	shortNewsItem model.NewsItems,
	fullNewsItem model.FullNewsItem) model.News {

	DBItem := model.News{}

	DBItem.Slug = fullNewsItem.Slug
	DBItem.Title = fullNewsItem.Title
	DBItem.Author = fullNewsItem.Author
	DBItem.PreviewText = fullNewsItem.PreviewText
	DBItem.Content = fullNewsItem.Content
	DBItem.ReadingTime = fullNewsItem.ReadingTime
	DBItem.PublishedAt = fullNewsItem.PublishedAt
	DBItem.Image = fullNewsItem.Image

	DBItem.SimilarNewsSlug = make([]string, len(fullNewsItem.SimilarNewsSlug))
	DBItem.PhotoReport = make([]string, len(fullNewsItem.PhotoReport))
	DBItem.TagsTitle = make([]string, len(shortNewsItem.TagsTitle))

	copy(DBItem.TagsTitle, shortNewsItem.TagsTitle)
	copy(DBItem.PhotoReport, fullNewsItem.PhotoReport)
	copy(DBItem.SimilarNewsSlug, fullNewsItem.SimilarNewsSlug)

	return DBItem
}
func (s *Service) ParseShortNews(ctx context.Context, limit int, offset int) ([]model.NewsItems, error) {
	resp, err := s.newsClient.GetNews(limit, offset)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Can`t parse news")
		return nil, err
	}

	NewsItems := make([]model.NewsItems, len(resp.Items), len(resp.Items))

	for i, elem := range resp.Items {
		item := model.NewsItems{}
		item.Slug = elem.Slug
		item.Title = elem.Title
		item.PreviewText = elem.PreviewText
		item.PublishedAt, err = s.tryParseTime(elem.PublishedAt.Day, elem.PublishedAt.Month, elem.PublishedAt.Year)
		item.ImagePreview = elem.ImagePreview

		for _, tag := range elem.Tags {
			item.TagsTitle = append(item.TagsTitle, tag.Title)
		}

		NewsItems[i] = item
	}

	return NewsItems, nil
}
func (s *Service) ParseFullNews(ctx context.Context, slug string) (model.FullNewsItem, error) {
	resp, err := s.newsClient.GetFullNews(slug)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Can`t parse full news")
		return model.FullNewsItem{}, err
	}

	FullNewsItems := model.FullNewsItem{}

	FullNewsItems.Slug = resp.Slug
	FullNewsItems.Title = resp.Title
	FullNewsItems.Author = resp.Author.Author
	FullNewsItems.PreviewText = resp.PreviewText
	FullNewsItems.Content = resp.Content
	FullNewsItems.ReadingTime = resp.ReadingTime
	FullNewsItems.PublishedAt, err = s.tryParseTime(resp.PublishedAt.Day, resp.PublishedAt.Month, resp.PublishedAt.Year)
	FullNewsItems.Image = resp.Image

	for _, photo := range resp.PhotoReport {
		FullNewsItems.PhotoReport = append(FullNewsItems.PhotoReport, photo.(string))
	}

	for _, similarNewsSlug := range resp.SimilarNews {
		FullNewsItems.SimilarNewsSlug = append(FullNewsItems.SimilarNewsSlug, similarNewsSlug.Slug)
	}

	return FullNewsItems, nil
}
func (s *Service) WriteDBNews(ctx context.Context, newsItem model.News) error {
	return s.repo.RewriteDBNews(ctx, newsItem)
}
func (s *Service) ReadDBNews(ctx context.Context, date time.Time) []model.News {
	return nil
}

func (s *Service) tryParseTime(d string, m string, y string) (time.Time, error) {
	year, err := strconv.Atoi(y)
	if err != nil {
		return time.Time{}, err
	}

	day, err := strconv.Atoi(d)
	if err != nil {
		return time.Time{}, err
	}

	month, exists := monthFormat[m]
	if !exists {
		return time.Time{}, fmt.Errorf("can`t search such moth format")
	}

	return time.Date(year, month, day, 0, 0, 0, 0, time.Local), nil
}
