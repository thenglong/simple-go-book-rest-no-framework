package entities

type Book struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Language      string `json:"language"`
	TotalPages    int    `json:"total_pages"`
	ImagePath     string `json:"image_path"`
	WikipediaLink string `json:"wikipedia_link"`
	Country       string `json:"country"`
}
