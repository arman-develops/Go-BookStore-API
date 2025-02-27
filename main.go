package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-book-api/config"
	"go-book-api/controllers"
	"go-book-api/models"
	"go-book-api/repositories"
	"go-book-api/routes"
)

func main() {

	//load configuration
	cfg := config.LoadConfig()

	//connect to database
	db, err := gorm.Open(postgres.Open(cfg.PostgresURI()), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	//automigrate the schema
	db.AutoMigrate(&models.Book{})

	//initialize repositories
	bookRepo := repositories.NewBookRepository(db)

	//initia controllers
	bookController := controllers.NewBookController(bookRepo)

	//setup routes
	router := routes.SetupRouter(bookController)

	//start server
	log.Printf("Server running on port %s", cfg.DBPort)
	router.Run(":" + cfg.ServerPort)
}
