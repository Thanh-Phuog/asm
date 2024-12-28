package main

import (
	"book_mana/database"
	"book_mana/routes"
)

func main() {
	database.Connect()

	router := routes.SetupRouter()

	router.Run(":8000")
}
