package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Kết nối đến PostgreSQL
func Connect() {

	dns := "user=myuser password=mypassword dbname=bookmana host=postgres port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"

	// dns := "host=127.0.0.1 user=myuser password=mypassword dbname=bookmana port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	var err error
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//// Migration
	// DB.AutoMigrate(&models.Book{})
	log.Println("Database connected successfully!")
}
