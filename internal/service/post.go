package service

import (
	"gorm.io/gorm"

	"github.com/Dominux/clean_architecture_blog/internal/entity"
	"github.com/Dominux/clean_architecture_blog/internal/repository"
)


type PostService struct {
	Entity entity.Post
	Repo repository.PostRepository
	DB gorm.DB
}

func (ps *PostService) CreatePost(title string, text string, author entity.User) (entity.Post, error) {
	post := entity.CreatePost(title, text, author)
	postRepository := repository.SavePost(post, ps.DB)
	return post, nil
}
