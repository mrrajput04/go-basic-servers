package main

import (
	"fmt"
	"simpleform/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber application instance
	app := fiber.New()

	// Define a health check endpoint to verify if the API is running
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("API-HEALTH: Healthly")
	})

	// Set up the application routes using a function from the routes package
	routes.Route(app)

	// Print a message to indicate the server is running
	fmt.Println("Server is Up!")
	// Start the Fiber web server on port 8081
	app.Listen(":8081")
}
