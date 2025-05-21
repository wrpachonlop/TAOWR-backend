package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/wrpachonlop/tool-inventory-backend/internal/db"
	"github.com/wrpachonlop/tool-inventory-backend/internal/routes"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Error loading .env file (local dev only)")
		}
	}

	db.ConnectDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.RegisterRoutes(r)
	r.Run(":8080")

}
