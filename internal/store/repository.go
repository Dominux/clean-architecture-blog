package store

import "github.com/Dominux/clean_architecture_blog/internal/entities"

type PostRepository interface {
	Create(p *entities.Post) (uint, error)
	Get() ([]*entities.Post, error)
	GetByID(id uint) (*entities.Post, error)
	Update(id uint, values map[string]interface{}) (*entities.Post, error)
	Delete(id uint) error
}

type CommentRepository interface {
	Create(c *entities.Comment, postID uint) (uint, error)
}
