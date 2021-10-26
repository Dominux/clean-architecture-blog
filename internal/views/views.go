package views

import "github.com/gin-gonic/gin"

type Views interface {
	Post() PostViews
	// Comment() CommentViews
}

type PostViews interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	GetObject(c *gin.Context)
}

type CommentViews interface {
	Create()
}
