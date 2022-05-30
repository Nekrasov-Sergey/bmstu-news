package model

type News struct {
	Slug   string `json:"slug"`
	Title  string `json:"title"`
	Author struct {
		OfferFrom interface{} `json:"offer_from"`
		Union     interface{} `json:"union"`
		Author    string      `json:"author"`
	} `json:"author"`
	PreviewText string `json:"preview_text"`
	Content     string `json:"content"`
	ReadingTime string `json:"reading_time"`
	PublishedAt struct {
		Day   string `json:"day"`
		Month string `json:"month"`
		Year  string `json:"year"`
	} `json:"published_at"`
	Image       string        `json:"image"`
	Views       int           `json:"views"` //не нужен
	PhotoReport []interface{} `json:"photoReport"`
	Documents   []interface{} `json:"documents"`
	SimilarNews []struct {
		Slug        string `json:"slug"`
		Title       string `json:"title"`
		PreviewText string `json:"preview_text"`
		PublishedAt struct {
			Day   string `json:"day"`
			Month string `json:"month"`
			Year  string `json:"year"`
		} `json:"published_at"`
		ImagePreview string `json:"imagePreview"`
		Tags         []struct {
			Id    int    `json:"id"`
			Slug  string `json:"slug"`
			Title string `json:"title"`
			Color string `json:"color"`
		} `json:"tags"`
	} `json:"similar_news"`
}

type AllNews struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Items  []struct {
		Slug        string `json:"slug"`
		Title       string `json:"title"`
		PreviewText string `json:"preview_text"`
		PublishedAt struct {
			Day   string `json:"day"`
			Month string `json:"month"`
			Year  string `json:"year"`
		} `json:"published_at"`
		ImagePreview string `json:"imagePreview"`
		Tags         []struct {
			Id    int    `json:"id"`
			Slug  string `json:"slug"`
			Title string `json:"title"`
			Color string `json:"color"`
		} `json:"tags"`
	} `json:"items"`
}

//общая структура из двух - хранится в базе данных
//посмотреть на видео - физически через gorm