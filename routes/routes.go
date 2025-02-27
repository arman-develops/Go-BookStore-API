package routes

import (
	"go-book-api/controllers"

	"github.com/gin-gonic/gin"

	"go-book-api/middleware"
)

func SetupRouter(bookController *controllers.BookController) *gin.Engine {
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

	}
	return router
}
