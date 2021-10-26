package web

import (
	"errors"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": id})
}

func (pv *PostViews) List(c *gin.Context) {
	posts, err := pv.service.Get()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": posts})
}

func (pv *PostViews) Receive(c *gin.Context) {
	rawID := c.Param("id")
	id, err := strconv.ParseUint(rawID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := pv.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, pv.service.Store.IsErrorNotFound()) {
			// Returing 404
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			// Returning 400
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": post})
}

func (pv *PostViews) Update(c *gin.Context) {
	rawID := c.Query("id")
	if rawID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id query param is not provided"})
		return
	}

	id, err := strconv.ParseUint(rawID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := pv.service.Update(uint(id), map[string]interface{}{
		"title":  c.PostForm("title"),
		"text":   c.PostForm("text"),
		"author": c.PostForm("author"),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": post})
}
