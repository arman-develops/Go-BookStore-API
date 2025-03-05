package controllers

import (
	"net/http"

	"go-book-api/models"
	"go-book-api/repositories"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	repo *repositories.AuthorRepository
}

// creates a new author controller
func NewAuthorController(repo *repositories.AuthorRepository) *AuthorController {
	return &AuthorController{repo: repo}
}

func (c *AuthorController) GetAuthors(ctx *gin.Context) {
	authors, err := c.repo.GetAllAuthors()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, authors)
}

func (c *AuthorController) GetAuthor(ctx *gin.Context) {
	id := ctx.Param("id")
	author, err := c.repo.GetAuthorWithID(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, author)
}

func (c *AuthorController) CreateAuthor(ctx *gin.Context) {
	var newAuthor models.Author
	if err := ctx.ShouldBindBodyWithJSON(&newAuthor); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAuthor, err := c.repo.CreateAuthor(newAuthor)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, createdAuthor)
}

func (c *AuthorController) UpdateAuthor(ctx *gin.Context) {
	id := ctx.Param("id")

	var author models.Author
	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author.AuthorID = id

	updatedAuthor, err := c.repo.UpdateAuthor(author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedAuthor)
}

// DeleteBook removes a book
func (c *AuthorController) DeleteAuthor(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.repo.DeleteAuthor(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
