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

func (pr *PostRepository) Get() ([]*entities.Post, error) {
	var dbPosts []PostModel

	// Getting results
	if result := pr.store.db.Find(&dbPosts); result.Error != nil {
		log.Printf("Error on creating post in database: %v", result.Error)
		return nil, result.Error
	}

	// Converting from []PostModel to []*entities.Post
	posts := make([]*entities.Post, len(dbPosts))
	for _, dbPost := range dbPosts {
		posts = append(posts, &entities.Post{
			Title: dbPost.Title,
			Text: dbPost.Text,
			Author: dbPost.Author,
		})
	}

	return posts, nil
}
