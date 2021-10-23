package gormstore

import (
	"log"

	"gorm.io/gorm"

	"github.com/Dominux/clean_architecture_blog/internal/entities"
)

type PostModel struct {
	gorm.Model
	Title  string
	Text   string
	Author *entities.User
}

type PostRepository struct {
	store *Store
}

func (pr *PostRepository) Create(p *entities.Post) (uint, error) {
	obj := PostModel{
		Title:  p.Title,
		Text:   p.Text,
		Author: p.Author,
	}

	if db := pr.store.db.Create(&obj); db.Error != nil {
		log.Printf("Error on creating post in database: %v", db.Error)
		return 0, db.Error
	}
	return obj.ID, nil
}
