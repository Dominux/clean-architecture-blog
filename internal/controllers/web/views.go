package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Dominux/clean_architecture_blog/internal/services"
)

type PostView struct {
	postService *services.PostService
}

func NewPostView(ps *services.PostService) *PostView {
	return &PostView{postService: ps}
}

func (pv *PostView) CreatePost(c *gin.Context) {
	title := c.PostForm("title")
	text := c.PostForm("text")
	author := c.PostForm("author")

	id, err := pv.postService.CreatePost(title, text, author)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusCreated, gin.H{"message": id})
}
