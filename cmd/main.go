package main

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"

	"github.com/Dominux/clean_architecture_blog/internal/store"
	"github.com/Dominux/clean_architecture_blog/internal/store/gormstore"
	"github.com/Dominux/clean_architecture_blog/internal/views/web"
)

const static = "frontend/build"

func main() {
	// Creating db
	db, err := store.NewDB(sqlite.Open("lol.db"))
	if err != nil {
		return
	}

	// Creating store
	store := gormstore.New(db)
	store.MigrateSchema()

	// Creating views
	webViews := web.New(store)

	// Creating router and filling it
	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob(static + "/*.html")
	router.Static("/static", static+"/static")

	router.Use(func(c *gin.Context) {
		// Pass api requests
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Next()
			return
		}

		c.HTML(http.StatusOK, "index.html", nil)
	})

	api_router := router.Group("/api")
	{
		postsRouter := api_router.Group("/posts")
		{
			postsRouter.POST("/", webViews.Post().Create)
			postsRouter.GET("/", webViews.Post().List)
			postsRouter.GET("/:id", webViews.Post().Receive)
			postsRouter.PUT("/:id", webViews.Post().Update)
			postsRouter.DELETE("/:id", webViews.Post().Delete)
		}
	}

	// Starting router
	router.Run()
}
