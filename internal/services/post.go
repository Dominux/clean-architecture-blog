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

	post := entities.NewPost(title, text, &a)

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
