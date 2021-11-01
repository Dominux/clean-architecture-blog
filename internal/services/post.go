package services

import (
	"github.com/Dominux/clean_architecture_blog/internal/entities"
	"github.com/Dominux/clean_architecture_blog/internal/store"
)

type PostService struct {
	Store store.Store
}

func NewPostService(s store.Store) *PostService {
	return &PostService{Store: s}
}

// Create Post, return its ID or error
func (ps *PostService) Create(title string, text string, author string) (uint, error) {
	// Conversion
	a := entities.User(author)

	post := &entities.Post{
		Title:  title,
		Text:   text,
		Author: &a,
	}

	postRepositoryID, err := ps.Store.Post().Create(post)
	if err != nil {
		return 0, err
	}

	return postRepositoryID, nil
}

func (ps *PostService) Get() ([]*entities.Post, error) {
	return ps.Store.Post().Get()
}

func (ps *PostService) GetByID(id uint) (*entities.Post, error) {
	return ps.Store.Post().GetByID(id)
}

func (ps *PostService) Update(id uint, values map[string]interface{}) (*entities.Post, error) {
	return ps.Store.Post().Update(id, values)
}

func (ps *PostService) Delete(id uint) error {
	return ps.Store.Post().Delete(id)
}
