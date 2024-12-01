package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"KSI-BE/config"
	"KSI-BE/routes"
)

func main() {
	// Initialize MongoDB
	config.Init()

	// Initialize Fiber app
	app := fiber.New()

	// Use CORS middleware with default settings
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",  // Allow all origins
		AllowMethods: "GET,POST,PUT,DELETE",  // Allow methods
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",  // Allow headers
	}))

	// Set up routes
	routes.Setup(app)

	// Start the app
	log.Fatal(app.Listen(":8080"))
}
