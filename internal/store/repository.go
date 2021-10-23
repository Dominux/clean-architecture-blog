package store

import "github.com/Dominux/clean_architecture_blog/internal/entities"

type PostRepository interface {
	Create(p *entities.Post) (uint, error)
}

type CommentRepository interface {
	Create(c *entities.Comment, postID uint) (uint, error)
}
