package main

import (
	"log"

	"KSI-BE/config"
	"KSI-BE/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize MongoDB
	config.Init()

	// Initialize Fiber app
	app := fiber.New()

	// Use CORS middleware with default settings
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                           // Allow all origins
		AllowMethods: "GET,POST,PUT,DELETE",                         // Allow methods
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", // Allow headers
	}))

	// Set up routes
	routes.Setup(app)

	// Start the app
	log.Fatal(app.Listen(":8080"))
}
