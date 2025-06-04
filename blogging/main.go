package main

import (
	"log"

	"blogging/routes"
	"blogging/services"

	// sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	sqladapter "github.com/Blank-Xu/sql-adapter"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/gofiber/swagger"
	"github.com/subosito/gotenv"
)

// @title Go Basecode API
// @version 1.0
// @description This is a sample API server with Casbin authentication and database connection.
// @host localhost:3000
// @BasePath /
func main() {
	gotenv.Load()
	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10 MB
	})
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080", // your React frontend URL
		AllowCredentials: true,                    // Allow credentials (cookies)
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization",
	}))

	// Initialize database
	db := services.InitDatabase()

	defer db.Close()

	_, err := sqladapter.NewAdapter(db, "sqlserver", "users")
	if err != nil {
		// handle error
	}
	if err != nil {
		panic(err)
	}

	// app.Use(func(c *fiber.Ctx) error {
	// 	c.Set("Access-Control-Allow-Origin", "*")
	// 	c.Set("Access-Control-Allow-Credentials", "true")
	// 	c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	// 	c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// 	if c.Method() == "OPTIONS" {
	// 		return c.SendStatus(200)
	// 	}
	// 	return c.Next()
	// })

	minio, err := services.InitMinio()
	if err != nil {
		log.Fatal("Failed to initialize MinIO client")
	}

	routes.RegisterAllRoutes(app, db, minio)

	// Swagger route
	// app.Get("/swagger/*", swagger.HandlerDefault)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
