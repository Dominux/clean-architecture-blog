package repository

import (
	"gorm.io/gorm"

	"github.com/Dominux/clean_architecture_blog/internal/entity"
)


type CommentRepository struct {
	gorm.Model
	entity.Comment
	PostID uint
}

func SaveComment(c entity.Comment, postID uint, db gorm.DB) CommentRepository {
	obj := CommentRepository{ 
		Author: c.Author, 
		Text: c.Text,
		PostID: postID,
	}
	db.Create(&obj)
	return obj
}

