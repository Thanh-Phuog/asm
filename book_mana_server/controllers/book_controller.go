package controllers

import (
	"book_mana/database"
	"book_mana/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello route for testing purposes
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, welcome to the Book API!"})
}

// Lấy tất cả sách
func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

// Tìm kiêm
func SearchBooks(c *gin.Context) {
	var books []models.Book
	if title := c.Query("title"); title != "" {
		database.DB.Where("title = ?", title).Find(&books)
		c.JSON(http.StatusOK, books)
	}

}

// Thêm sách mới
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&book)
	c.JSON(http.StatusOK, book)
}

// Cập nhật sách
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

// Xóa sách
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	database.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
