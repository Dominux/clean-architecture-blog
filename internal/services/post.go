package services

import (
	"gorm.io/gorm"

	"github.com/Dominux/clean_architecture_blog/internal/entities"
	"github.com/Dominux/clean_architecture_blog/internal/repository"
)

type PostService struct {
	DB *gorm.DB
}

// Create Post, return its ID or error
func (ps *PostService) CreatePost(title string, text string, author str) (uint, error) {
	// Converting author to the User type
	author := new(*entities.User, author)

	post := entities.CreatePost(title, text, author)

	postRepository, err := repository.SavePost(post, *ps.DB)
	if err != nil {
		return 0, err
	}

	return postRepository.ID, nil
}
