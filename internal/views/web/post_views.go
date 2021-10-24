package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Dominux/clean_architecture_blog/internal/services"
)

type PostViews struct {
	service *services.PostService
}

func (pv *PostViews) Create(c *gin.Context) {
	title := c.PostForm("title")
	text := c.PostForm("text")
	author := c.PostForm("author")

	id, err := pv.service.Create(title, text, author)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusCreated, gin.H{"message": id})
}

func (pv *PostViews) Get(c *gin.Context) {
	posts, err := pv.service.Get()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": posts})
}
