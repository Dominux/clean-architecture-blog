package store

type Store interface {
	Post() PostRepository
	IsErrorNotFound() error
	// Comment() CommentRepository
}
