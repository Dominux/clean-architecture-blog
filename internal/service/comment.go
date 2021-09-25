package service

import (
	"gorm.io/gorm"

	"github.com/Dominux/clean_architecture_blog/internal/entity"
	"github.com/Dominux/clean_architecture_blog/internal/repository"
)


type CommentService struct {
	DB *gorm.DB
}

func (cs *CommentService) CreateComment(postID uint, author *entity.User, text string) (entity.Comment, error) {
	// TODO: create binding the post to the comment when it's initializing
	comment := entity.CreateComment(nil, author, text)
	commentRepository := repository.SaveComment(comment, postID, *cs.DB)
	return comment, nil
}

