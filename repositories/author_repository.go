package repositories

import (
	"errors"
	"go-book-api/models"

	"gorm.io/gorm"
)

// handles database operations for author
type AuthorRepository struct {
	DB *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{DB: db}
}

func (r *AuthorRepository) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	result := r.DB.Find(&authors)
	return authors, result.Error
}

func (r *AuthorRepository) GetAuthorWithID(id string) (models.Author, error) {
	var author models.Author

	result := r.DB.First(&author, id)
	return author, result.Error
}

func (r *AuthorRepository) CreateAuthor(author models.Author) (models.Author, error) {
	result := r.DB.Create(&author)
	return author, result.Error
}

func (r *AuthorRepository) UpdateAuthor(author models.Author) (models.Author, error) {
	var existingAuthor models.Author

	if err := r.DB.First(&existingAuthor, author.AuthorID).Error; err != nil {
		return author, errors.New("Author Not found")
	}
	result := r.DB.Save(&author)
	return author, result.Error
}

func (r *AuthorRepository) DeleteAuthor(id string) error {
	result := r.DB.Delete(&models.Book{}, id)
	if result.RowsAffected == 0 {
		return errors.New("Author Not found")
	}
	return result.Error
}
