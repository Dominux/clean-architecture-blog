package store

type Store interface {
	Post() PostRepository
	// Comment() CommentRepository
}
