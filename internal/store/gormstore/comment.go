package gormstore

import (
	"log"

	"gorm.io/gorm"

	"github.com/Dominux/clean_architecture_blog/internal/entities"
)


type CommentModel struct {
	gorm.Model
	entities.Comment
	PostID uint
}


type CommentRepository struct {
	store *Store
}

func (cr *CommentRepository) Create(c entities.Comment, postID uint, db gorm.DB) (uint, error) {
	obj := CommentModel{ 
		Author: c.Author, 
		Text: c.Text,
		PostID: postID,
	}

	if db = *cr.store.db.Create(&obj); db.Error != nil {
		log.Printf("Error on creating post in database: %v", db.Error)
		return 0, db.Error
	}
	return obj.ID, nil
}

