package model

type DBNews struct {
	Slug             string   `json:"slug"`               //Slug новости (переход)
	Title            string   `json:"title"`              //Название новости
	Author           string   `json:"author"`             //Автор новости - оставили одно поле из структуры
	PreviewText      string   `json:"preview_text"`       //Иногд а совпадает с Title, иногда с Content, есть малые отличия
	Content          string   `json:"content"`            //Полный текст новости
	ReadingTime      string   `json:"reading_time"`       //Время прочтения
	PublishedAtDay   int      `json:"published_at_day"`   //Дата публикации
	PublishedAtMonth string   `json:"published_at_month"` //Дата публикации
	PublishedAtYear  int      `json:"published_at_year"`  //Дата публикации
	Image            string   `json:"image"`              //Первая картинка новости
	PhotoReport      []string `gorm:"type:text[]"`        //Несколько фотографий для пролистывания
	SimilarNewsSlug  []string `gorm:"type:text[]"`        //Похожие новости
	TagsTitle        []string `gorm:"type:text[]"`        //Тэги для фильтрации
}

type FullNewsItems struct {
	Slug             string   `json:"slug"`               //Slug новости (переход)
	Title            string   `json:"title"`              //Название новости
	Author           string   `json:"author"`             //Автор новости - оставили одно поле из структуры
	PreviewText      string   `json:"preview_text"`       //Иногд а совпадает с Title, иногда с Content, есть малые отличия
	Content          string   `json:"content"`            //Полный текст новости
	ReadingTime      string   `json:"reading_time"`       //Время прочтения
	PublishedAtDay   int      `json:"published_at_day"`   //Дата публикации
	PublishedAtMonth int      `json:"published_at_month"` //Дата публикации
	PublishedAtYear  int      `json:"published_at_year"`  //Дата публикации
	Image            string   `json:"image"`              //Первая картинка новости
	Views            int      `json:"views"`              //Бесполезно
	PhotoReport      []string `json:"photoReport"`        //Несколько фотографий для пролистывания
	SimilarNewsSlug  []string `json:"similar_news_slug"`  //Похожие новости
}

//NewsItems - краткая новость
type NewsItems []struct {
	Slug             string   `json:"slug"`               //Нужен для перехода к полной новости
	Title            string   `json:"title"`              //Назавание новости
	PreviewText      string   `json:"preview_text"`       //Не показывается
	PublishedAtDay   int      `json:"published_at_day"`   //Дата публикации
	PublishedAtMonth int      `json:"published_at_month"` //Дата публикации
	PublishedAtYear  int      `json:"published_at_year"`  //Дата публикации
	ImagePreview     string   `json:"imagePreview"`       //Превью картинка
	TagsTitle        []string `json:"tags_title"`         //Тэги для фильтрации
}

//общая структура из двух - хранится в базе данных
//посмотреть на видео - физически через gorm
