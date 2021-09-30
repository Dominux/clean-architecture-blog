package entities


type Comment struct {
	Post *Post
	Author *User	
	Text string
}

func CreateComment(p *Post, a *User, t string) Comment {
	return Comment{
		Post: p,
		Author: a,
		Text: t,
	}
}
