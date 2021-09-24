package entity


type Post struct {
	Title string
	Text string
	Author User
	Comments []*Comment
}

func CreatePost(title string, text string, author User) Post {
	return Post{
		Title: title,
		Text: text,
		Author: author,
	}
}
