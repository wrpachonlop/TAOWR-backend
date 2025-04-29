package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	// Configure GORM settings
	config := &gorm.Config{
		PrepareStmt: false, // 👈 important!
		Logger:      logger.Default.LogMode(logger.Info),
	}

	// Open connection
	database, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Get the generic DB object sql.DB to set connection pool options
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("Failed to get generic database object: %v", err)
	}

	// Recommended settings for Supabase / PostgreSQL
	sqlDB.SetMaxIdleConns(10)                  // Max idle connections
	sqlDB.SetMaxOpenConns(100)                 // Max open connections
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Reset connections every 30 min

	DB = database
	log.Println("Database connected successfully!")
}
