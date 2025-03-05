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

	//automigrate the schemas
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Author{})

	//initialize repositories
	bookRepo := repositories.NewBookRepository(db)
	authorRepo := repositories.NewAuthorRepository(db)

	//initia controllers
	bookController := controllers.NewBookController(bookRepo)
	authorController := controllers.NewAuthorController(authorRepo)

	//setup routes
	router := routes.SetupRouter(bookController, authorController)

	//start server
	log.Printf("Server running on port %s", cfg.DBPort)
	router.Run(":" + cfg.ServerPort)
}
