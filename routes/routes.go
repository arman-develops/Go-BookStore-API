package routes

import (
	"go-book-api/controllers"

	"github.com/gin-gonic/gin"

	"go-book-api/middleware"
)

func SetupRouter(
	bookController *controllers.BookController,
	authorController *controllers.AuthorController,
) *gin.Engine {
	router := gin.Default()

	//apply global middleware
	router.Use(middleware.Logger())

	//v2 API routes
	v2 := router.Group("/api/v2")
	{
		books := v2.Group("/books")
		{
			books.GET("", bookController.GetAllBooks)
			books.GET("/:id", bookController.GetBook)
			books.POST("", bookController.CreateBook)
			books.PUT("/:id", bookController.UpdateBook)
			books.DELETE("/:id", bookController.DeleteBook)
		}

		authors := v2.Group("/authors")
		{
			authors.GET("", authorController.GetAuthors)
			authors.GET("/:id", authorController.GetAuthor)
			authors.POST("", authorController.CreateAuthor)
			authors.PUT("/:id", authorController.UpdateAuthor)
			authors.DELETE("/:id", authorController.DeleteAuthor)
		}

	}

	return router
}
