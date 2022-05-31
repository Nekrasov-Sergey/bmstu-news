package news

type ResponseFullNews struct {
	Slug   string   `json:"slug"`  //Slug новости (переход)
	Title  string   `json:"title"` //Название новости
	Author struct { //Автор публикации
		OfferFrom interface{} `json:"offer_from"` //Всегда null
		Union     interface{} `json:"union"`      //Всегда null
		Author    string      `json:"author"`
	} `json:"author"`
	PreviewText string   `json:"preview_text"` //Иногда совпадает с Title, иногда с Content, есть малые отличия
	Content     string   `json:"content"`      //Полный текст новости
	ReadingTime string   `json:"reading_time"` //Время прочтения
	PublishedAt struct { //Дата публикации
		Day   string `json:"day"`
		Month string `json:"month"`
		Year  string `json:"year"`
	} `json:"published_at"`
	Image       string        `json:"image"`       //Первая картинка новости
	Views       int           `json:"views"`       //Бесполезно
	PhotoReport []interface{} `json:"photoReport"` //Несколько фотографий для пролистывания
	Documents   []interface{} `json:"documents"`   //Не нашел новость, где не пустой
	SimilarNews []struct {    //Похожие новости
		Slug        string   `json:"slug"`         //Slug новости
		Title       string   `json:"title"`        //Название
		PreviewText string   `json:"preview_text"` //Не используется, пока оставим
		PublishedAt struct { //Дата публикации
			Day   string `json:"day"`
			Month string `json:"month"`
			Year  string `json:"year"`
		} `json:"published_at"`
		ImagePreview string     `json:"imagePreview"` //Превью картинка
		Tags         []struct { //Тэги для фильтрации
			Id    int    `json:"id"`    //ID тэга
			Slug  string `json:"slug"`  //Slug тэга
			Title string `json:"title"` //Название тэга
			Color string `json:"color"` //Цвет тэга
		} `json:"tags"`
	} `json:"similar_news"`
}
