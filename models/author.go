package models

type Author struct {
	AuthorID   string `json:"author_id" gorm:"primaryKey"`
	AuthorName string `json:"author_name" gorm:"size:255;not null"`
}
