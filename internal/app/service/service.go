package service

import (
	"context"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/model"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/app/repository"
	"github.com/Nekrasov-Sergey/bmstu-news.git/internal/pkg/clients/news"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

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
		/*item.PublishedAtDay, _ = strconv.Atoi(elem.PublishedAt.Day)
		item.PublishedAtMonth = elem.PublishedAt.Month
		item.PublishedAtYear, _ = strconv.Atoi(elem.PublishedAt.Year)*/
		item.PublishedAt, err = s.tryParseTime(elem.PublishedAt.Day, elem.PublishedAt.Month, elem.PublishedAt.Year)
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
	FullNewsItems.PublishedAt, err = s.tryParseTime(resp.PublishedAt.Day, resp.PublishedAt.Month, resp.PublishedAt.Year)
	/*FullNewsItems.PublishedAtDay, _ = strconv.Atoi(resp.PublishedAt.Day)
	FullNewsItems.PublishedAtMonth = resp.PublishedAt.Month
	FullNewsItems.PublishedAtYear, _ = strconv.Atoi(resp.PublishedAt.Year)*/
	FullNewsItems.Image = resp.Image

	for _, photo := range resp.PhotoReport {
		FullNewsItems.PhotoReport = append(FullNewsItems.PhotoReport, photo.(string))
	}

	for _, similarNewsSlug := range resp.SimilarNews {
		FullNewsItems.SimilarNewsSlug = append(FullNewsItems.SimilarNewsSlug, similarNewsSlug.Slug)
	}

	return FullNewsItems, nil
}
func (s *Service) WriteDBNews(ctx context.Context, NewsItems model.NewsItems, FullNewsItems model.FullNewsItems) error {
	//сделать тут соединение двух структур в DBNews и отправить эту структуру в repo.RewriteDBNews(ctx context.Context, DBitems model.DBNews)

	DBItems := model.DBNews{}

	DBItems.Slug = FullNewsItems.Slug
	DBItems.Title = FullNewsItems.Title
	DBItems.Author = FullNewsItems.Author
	DBItems.PreviewText = FullNewsItems.PreviewText
	DBItems.Content = FullNewsItems.Content
	DBItems.ReadingTime = FullNewsItems.ReadingTime
	DBItems.PublishedAt = FullNewsItems.PublishedAt
	DBItems.Image = FullNewsItems.Image

	for _, tag := range NewsItems.TagsTitle {
		DBItems.TagsTitle = append(DBItems.TagsTitle, tag)
	}

	for _, photo := range FullNewsItems.PhotoReport {
		DBItems.PhotoReport = append(DBItems.PhotoReport, photo)
	}

	for _, similarNewsSlug := range FullNewsItems.PhotoReport {
		DBItems.SimilarNewsSlug = append(DBItems.SimilarNewsSlug, similarNewsSlug)
	}

	return s.repo.RewriteDBNews(ctx, DBItems)
}
func (s *Service) ReadDBNews(ctx context.Context, date time.Time) []model.DBNews {
	return nil
}

func (s *Service) tryParseTime(d string, m string, y string) (time.Time, error) {

	//year, _ := strconv.Atoi(y)
	day, _ := strconv.Atoi(d)

	/*monthFormat := map[string]time.Time{
		"января":   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		"февраля":  time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
		"марта":    time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC),
		"апреля":   time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
		"мая":      time.Date(2022, 5, 1, 0, 0, 0, 0, time.UTC),
		"июня":     time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC),
		"июля":     time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC),
		"августа":  time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
		"сентября": time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC),
		"октября":  time.Date(2022, 10, 1, 0, 0, 0, 0, time.UTC),
		"ноября":   time.Date(2022, 11, 1, 0, 0, 0, 0, time.UTC),
		"декабря":  time.Date(2022, 12, 1, 0, 0, 0, 0, time.UTC),
	}*/

	if day < 10 {
		d = "0" + d
	}

	//date := time.Date(year, monthFormat[m].Month(), day, 0, 0, 0, 0, time.Local)
	//return date, nil

	monthFormat := map[string]string{
		"января":   "01",
		"февраля":  "02",
		"марта":    "03",
		"апреля":   "04",
		"мая":      "05",
		"июня":     "06",
		"июля":     "07",
		"августа":  "08",
		"сентября": "09",
		"октября":  "10",
		"ноября":   "11",
		"декабря":  "12",
	}

	date := y + "-" + monthFormat[m] + "-" + d
	res, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}

	return res, nil
}
