package views

type Views interface {
	Post() *PostViews
	Comment() *CommentViews
}

type PostViews interface {
	Create()
}

type CommentViews interface {
	Create()
}
