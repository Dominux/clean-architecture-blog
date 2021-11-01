package entities

type Post struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author *User  `json:"author"`
}
