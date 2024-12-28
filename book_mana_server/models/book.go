package models

import "gorm.io/gorm"

// Book struct
type Book struct {
	gorm.Model
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   string `json:"year"`
}
