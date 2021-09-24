package repository

import (
	"gorm.io/gorm"

	"github.com/Dominux/clean_architecture_blog/internal/entity"
)


type PostRepository struct {
	gorm.Model
	entity.Blog
}

func SavePost(p entity.Post, db gorm.DB) PostRepository {
	obj := PostRepository{
		Title: p.Title,
		Text: p.Text,
		Author: p.Author,
	}
	db.Create(&obj)
	return obj
}
