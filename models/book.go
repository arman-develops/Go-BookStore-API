package models

// represents information about the book
type Book struct {
	BookID      string `json:"id" gorm:"primaryKey`
	AuthorID    string `json:"author_id" gorm:"size:255;not null;unique`
	Title       string `json:"title" gorm:"size:255;not null"`
	PublishDate string `json:"publish_date" gorm:size:255;not null`
	Description string `json:"description" gorm:"size:255; not null"`
	Genre       string `json:"genre" gorm:"size:255; not null"`
}
