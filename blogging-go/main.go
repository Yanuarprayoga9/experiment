package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/denisenkom/go-mssqldb" // SQL Server driver
	
	"myblog/internal/app/post" // Adjust import path
)

func main() {
	// Database connection
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// SQL Server connection string format
		dbURL = "sqlserver://Sa:YourStrongPassword123!@localhost:1433?database=blogging"
	}

	db, err := sqlx.Connect("sqlserver", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Create Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	// API v1 group
	v1 := router.Group("/api/v1")
	
	// Register post routes
	post.RegisterPostRoutes(v1, db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}