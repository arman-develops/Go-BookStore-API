package controllers

import (
	"net/http"

	"go-book-api/models"
	"go-book-api/repositories"

	"github.com/gin-gonic/gin"
)

// book controller that handles requests related to books
type BookController struct {
	repo *repositories.BookRepository
}

// creates a new book controller
func NewBookController(repo *repositories.BookRepository) *BookController {
	return &BookController{repo: repo}
}

// list of all books
func (c *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := c.repo.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, books)
}

// get a single book
func (c *BookController) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	book, err := c.repo.GetBookByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	}

	ctx.IndentedJSON(http.StatusOK, book)
}

// create a new book: from json request
func (c *BookController) CreateBook(ctx *gin.Context) {
	var newBook models.Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBook, err := c.repo.CreateBook(newBook)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, createdBook)
}

// UpdateBook updates a book
func (c *BookController) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
	// 	return
	// }

	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.BookID = id

	updatedBook, err := c.repo.UpdateBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedBook)
}

// DeleteBook removes a book
func (c *BookController) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
	// 	return
	// }

	if err := c.repo.DeleteBook(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
