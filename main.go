package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// define your schemas and data structures
type Book struct {
	BookID      string `json:"bookID"`
	AuthorID    string `json:"authorID`
	Title       string `json:"title"`
	PublishDate string `json:"publish_date`
	Description string `json:"description"`
	Genre       string `json:"genre"`
}

type Author struct {
	AuthorID   string `json:"authorID"`
	AuthorName string `json:"author_name`
}

// data to be stored in memory
var books = []Book{
	{BookID: "1", AuthorID: "001", Title: "Angels and Demons", PublishDate: "2009/2/4", Description: "Awesome", Genre: "Thriller"},
	{BookID: "2", AuthorID: "002", Title: "Great Gatsby", PublishDate: "2014/11/2", Description: "Awesome 1", Genre: "Drama"},
	{BookID: "3", AuthorID: "001", Title: "Star Trek", PublishDate: "2017/8/7", Description: "Awesome 3", Genre: "Sci-fi"},
	{BookID: "4", AuthorID: "002", Title: "Prince of Argos", PublishDate: "2019/12/1", Description: "Awesome 4", Genre: "Medievial"},
	{BookID: "5", AuthorID: "002", Title: "Sands of Time", PublishDate: "2009/2/4", Description: "Awesome 5", Genre: "Box Office"},
}

var authors = []Author{
	{AuthorID: "001", AuthorName: "Horjben Jurl"},
	{AuthorID: "002", AuthorName: "Kent Robby"},
}

func main() {
	router := gin.Default()
	//first version of the books api works in memory
	//get all authors
	router.GET("/api/v1/authors", getAuthors)

	//get all books
	router.GET("/api/v1/books", getBooks)

	//add new author
	router.POST("/api/v1/authors", addAuthor)

	//add new book
	router.POST("/api/v1/books", addBook)

	//get author by id
	router.GET("/api/v1/authors/:id", getAuthorByID)

	//get book by id
	router.GET("/api/v1/books/:id", getBookByID)

	//get book by authorID
	router.GET("/api/v1/books/author/:authorID", getBookByAuthorID)

	//update book
	router.PUT("/api/v1/books/:id", updateBook)

	//update author
	router.PUT("/api/v1/authors/:id", updateAuthor)

	router.Run("localhost:3000")
}

// write code to handle the endpoints
// get all authors
func getAuthors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, authors)
}

// get all books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// add a new author
func addAuthor(c *gin.Context) {
	var newAuthor Author
	if err := c.BindJSON(&newAuthor); err != nil {
		return
	}

	authors = append(authors, newAuthor)
	c.IndentedJSON(http.StatusCreated, newAuthor)
}

// add a new book
func addBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// get all books by a certain author id
func getBookByAuthorID(c *gin.Context) {
	id := c.Param("authorID")

	for _, book := range books {
		if book.AuthorID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	//process a book not found error incase its not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book for author not found"})
}

// get a book with book id
func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.BookID == id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book with id not found"})

}

// get author with author id
func getAuthorByID(c *gin.Context) {
	id := c.Param("id")

	for _, author := range authors {
		if author.AuthorID == id {
			c.IndentedJSON(http.StatusOK, author)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Author by ID not found"})
}

// update author
func updateAuthor(c *gin.Context) {
	id := c.Param("id")
	var updatedAuthor Author

	if err := c.BindJSON(&updatedAuthor); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	for i, author := range authors {
		if author.AuthorID == id {
			updatedAuthor.AuthorID = id
			authors[i] = updatedAuthor
			c.IndentedJSON(http.StatusOK, updatedAuthor)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Author not found"})
}

// update book
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook Book

	if err := c.BindJSON(&updatedBook); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"Error": err.Error()})
	}

	for i, book := range books {
		if book.BookID == id {
			updatedBook.BookID = id
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}
