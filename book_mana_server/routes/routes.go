package routes

import (
	"book_mana/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//Cấu hình CORS
	router.Use(cors.Default())

	// Định nghĩa các route cho sách
	router.GET("/hello", controllers.Hello)
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
	router.GET("/books/search", controllers.SearchBooks)

	return router
}
