package service

import (
	"context"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/model"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/pkg/clients/news"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Service struct {
	newsClient *news.Client
}

func New(ctx context.Context) (*Service, error) {
	newsClient := news.New(ctx)

	return &Service{
		newsClient: newsClient,
	}, nil
}

func (s *Service) ParseNews(ctx context.Context, limit string, offset string) ([]model.NewsItems, error) {
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
		item.PublishedAtDay, _ = strconv.Atoi(elem.PublishedAt.Day)
		item.PublishedAtMonth = elem.PublishedAt.Month
		item.PublishedAtYear, _ = strconv.Atoi(elem.PublishedAt.Year)
		item.ImagePreview = elem.ImagePreview

		for _, tag := range elem.Tags {
			item.TagsTitle = append(item.TagsTitle, tag.Title)
		}

		NewsItems[i] = item
	}

	return NewsItems, nil
}
func (s *Service) ParseFullNews(ctx context.Context, slug string) (model.FullNewsItems, error) {
	resp, err := s.newsClient.GetFullNews(slug)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Can`t parse full news")
		return model.FullNewsItems{}, err
	}

	FullNewsItems := model.FullNewsItems{}

	FullNewsItems.Slug = resp.Slug
	FullNewsItems.Title = resp.Title
	FullNewsItems.Author = resp.Author.Author
	FullNewsItems.PreviewText = resp.PreviewText
	FullNewsItems.Content = resp.Content
	FullNewsItems.ReadingTime = resp.ReadingTime
	FullNewsItems.PublishedAtDay, _ = strconv.Atoi(resp.PublishedAt.Day)
	FullNewsItems.PublishedAtMonth = resp.PublishedAt.Month
	FullNewsItems.PublishedAtYear, _ = strconv.Atoi(resp.PublishedAt.Year)
	FullNewsItems.Image = resp.Image

	for _, photo := range resp.PhotoReport {
		FullNewsItems.PhotoReport = append(FullNewsItems.PhotoReport, photo.(string))
	}

	for _, similarNewsSlug := range resp.PhotoReport {
		FullNewsItems.SimilarNewsSlug = append(FullNewsItems.SimilarNewsSlug, similarNewsSlug.(string))
	}

	return FullNewsItems, nil
}
func (s *Service) WriteDBNews(ctx context.Context, NewsItems model.NewsItems, FullNewsItems model.FullNewsItems) error {
	//сделать тут соединение двух структур в DBNews и отправить эту структуру в repo.RewriteWorkingOutItems(ctx context.Context, items []model.DBNews)
	return nil
}
func (s *Service) ReadDBNews(ctx context.Context, date time.Time) []model.DBNews {
	return nil
}
