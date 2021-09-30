package services

import (
	"gorm.io/gorm"

	"github.com/Dominux/clean_architecture_blog/internal/entities"
	"github.com/Dominux/clean_architecture_blog/internal/repository"
)


type CommentService struct {
	DB *gorm.DB
}

// Create Comment, return its ID
func (cs *CommentService) CreateComment(postID uint, author *entities.User, text string) (uint, error) {
	comment := entities.CreateComment(nil, author, text)
	commentRepository := repository.SaveComment(comment, postID, *cs.DB)
	return commentRepository.ID, nil
}

