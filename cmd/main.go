package main

import (
	"log"
	"os"

	"github.com/Rafasanchez/wsapi/pkg/api/routes"
	"github.com/Rafasanchez/wsapi/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up the database connection
	db, err := config.SetupDB()
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	defer db.Close()

	// Create a new Gin router
	router := gin.Default()

	// Set up the API routes
	routes.SetupRoutes(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
