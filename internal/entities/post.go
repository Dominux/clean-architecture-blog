package entities

type Post struct {
	Title  string
	Text   string
	Author *User
}

func NewPost(title string, text string, author *User) *Post {
	return &Post{
		Title:  title,
		Text:   text,
		Author: author,
	}
}
