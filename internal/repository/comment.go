package repository

import (
	"gorm.io/gorm"

	"github.com/Dominux/clean_architecture_blog/internal/entity"
)


type CommentRepository struct {
	gorm.Model
	entity.Comment
}

func SaveComment(c entity.Comment, db gorm.DB) CommentRepository {
	obj := CommentRepository{ Author: c.Author, Text: c.Text }
	db.Create(&obj)
	return obj
}

