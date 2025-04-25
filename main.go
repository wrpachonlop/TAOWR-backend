package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/wrpachonlop/tool-inventory-backend/internal/db"
	"github.com/wrpachonlop/tool-inventory-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
