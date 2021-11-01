package entities

type Post struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author *User  `json:"author"`
}

func NewPost(title string, text string, author *User) *Post {
	return &Post{
		Title:  title,
		Text:   text,
		Author: author,
	}
}
