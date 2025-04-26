package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wrpachonlop/tool-inventory-backend/internal/db"
	"github.com/wrpachonlop/tool-inventory-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting the application...")
	// Load environment variables from .env file if not in production
	log.Println("Loading environment variables...")
	// Check if the ENV variable is set to production
	log.Println(("ENV: " + os.Getenv("ENV")))
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Error loading .env file (local dev only)")
		}
	}

	db.ConnectDB()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
