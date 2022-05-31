package news

type ResponseNews struct {
	Total  int        `json:"total"`  //Новостей всего
	Limit  int        `json:"limit"`  //Сколько новостей выдано
	Offset int        `json:"offset"` //Смещение по новостям
	Items  []struct { //Слайс кратких данных по новостям
		Slug        string   `json:"slug"`         //Нужен для перехода к полной новости
		Title       string   `json:"title"`        //Назавание новости
		PreviewText string   `json:"preview_text"` //Не показывается
		PublishedAt struct { //Дата публикации
			Day   string `json:"day"`
			Month string `json:"month"`
			Year  string `json:"year"`
		} `json:"published_at"`
		ImagePreview string     `json:"imagePreview"` //Превью картинка
		Tags         []struct { //Тэги для фильтрации
			Id    int    `json:"id"`    //ID тэга
			Slug  string `json:"slug"`  //Slug тэга, оставим
			Title string `json:"title"` //Название тэга
			Color string `json:"color"` //Цвет тэга
		} `json:"tags"`
	} `json:"items"`
}
