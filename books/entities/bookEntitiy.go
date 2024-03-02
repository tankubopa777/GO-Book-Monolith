package entities

type Book struct {
	ID            int64  `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	ISBN          string `json:"isbn"`
	PublishedDate string `json:"published_date"`
}
