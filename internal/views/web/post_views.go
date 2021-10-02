package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Dominux/clean_architecture_blog/internal/services"
)

type PostViews struct {
	views *Views
	postService *services.PostService
}

func NewPostViews(v *Views) *PostViews {
	return &PostViews{
		views: v,
		postService: services.NewPostService(v.store),
	}
}

func (pv *PostViews) Create(c *gin.Context) {
	title := c.PostForm("title")
	text := c.PostForm("text")
	author := c.PostForm("author")

	id, err := pv.postService.Create(title, text, author)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusCreated, gin.H{"message": id})
}
