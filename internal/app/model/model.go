package model

type News struct {
	FullNewsItems FullNewsItems
	NewsItems     NewsItems
}

type FullNewsItems struct {
	Slug        string        `json:"slug"`         //Slug новости (переход)
	Title       string        `json:"title"`        //Название новости
	Author      string        `json:"author"`       //Автор новости - оставили одно поле из структуры
	PreviewText string        `json:"preview_text"` //Иногд а совпадает с Title, иногда с Content, есть малые отличия
	Content     string        `json:"content"`      //Полный текст новости
	ReadingTime string        `json:"reading_time"` //Время прочтения
	PublishedAt PublishedAt   `json:"published_at"` //Дата публикации
	Image       string        `json:"image"`        //Первая картинка новости
	Views       int           `json:"views"`        //Бесполезно
	PhotoReport []interface{} `json:"photoReport"`  //Несколько фотографий для пролистывания
	Documents   []interface{} `json:"documents"`    //?? todo: найти новость, где documents не пустой
	SimilarNews SimilarNews   `json:"similar_news"` //Похожие новости
}

//NewsItems - краткая новость
type NewsItems []struct {
	Slug         string      `json:"slug"`         //Нужен для перехода к полной новости
	Title        string      `json:"title"`        //Назавание новости
	PreviewText  string      `json:"preview_text"` //Не показывается
	PublishedAt  PublishedAt `json:"published_at"` //Дата публикации
	ImagePreview string      `json:"imagePreview"` //Превью картинка
	Tags         Tags        `json:"tags"`         //Тэги для фильтрации
}

type SimilarNews []struct {
	Slug         string      `json:"slug"`         //Slug новости
	Title        string      `json:"title"`        //Название
	PreviewText  string      `json:"preview_text"` //Не используется, пока оставим
	PublishedAt  PublishedAt `json:"published_at"` //Дата публикации
	ImagePreview string      `json:"imagePreview"` //Превью картинка
	Tags         Tags        `json:"tags"`         //Тэги для фильтрации
}

type PublishedAt struct {
	Day   string `json:"day"`
	Month string `json:"month"`
	Year  string `json:"year"`
}

type Tags []struct {
	Id    int    `json:"id"`    //ID тэга
	Slug  string `json:"slug"`  //Slug тэга, оставим
	Title string `json:"title"` //Название тэга
	Color string `json:"color"` //Цвет тэга
}

//общая структура из двух - хранится в базе данных
//посмотреть на видео - физически через gorm
