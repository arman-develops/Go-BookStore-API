package repositories

import (
	"errors"
	"go-book-api/models"

	"gorm.io/gorm"
)

// handles database operations for books
type BookRepository struct {
	DB *gorm.DB
}

// creates a new repository
func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

// GetAllBooks: returns all books from the database
func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := r.DB.Find(&books)
	return books, result.Error
}

// GetBookByID: finds a book by ID
func (r *BookRepository) GetBookByID(id string) (models.Book, error) {
	var book models.Book
	result := r.DB.First(&book, id)
	return book, result.Error
}

// //GetBookByAuthorID: finds books by a specific author
// func (r *BookRepository) GetBookByAuthorID(id string) (models.Book, error) {
// 	var book models.Book
// 	result := r.DB.First(&book.AuthorID, id)
// 	return book, result.Error
// }

// CreateBook: adds a new book to the respository
func (r *BookRepository) CreateBook(book models.Book) (models.Book, error) {
	result := r.DB.Create(&book)
	return book, result.Error
}

func (r *BookRepository) UpdateBook(book models.Book) (models.Book, error) {
	var existingBook models.Book
	if err := r.DB.First(&existingBook, book.BookID).Error; err != nil {
		return book, errors.New("Book not found")
	}

	result := r.DB.Save(&book)
	return book, result.Error
}

func (r *BookRepository) DeleteBook(id string) error {
	result := r.DB.Delete(&models.Book{}, id)
	if result.RowsAffected == 0 {
		return errors.New("books not found")
	}
	return result.Error
}
