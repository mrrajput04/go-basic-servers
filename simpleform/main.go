package main

import (
	"fmt"
	"simpleform/internal/database"
	"simpleform/internal/routes"
	"simpleform/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber application instance
	app := fiber.New()

	// Define a health check endpoint to verify if the API is running
	app.Get("/health-check", func(c *fiber.Ctx) error {
		return c.SendString("API-HEALTH: Healthly")
	})

	// Initialize a database client using a custom function from the database package
	db := database.DBClient()
	// Use a middleware to attach the database client to the Fiber context, making it accessible in request handlers
	app.Use(utils.DbMiddleware(db))

	// Set up the application routes using a function from the routes package
	routes.Route(app)

	// Print a message to indicate the server is running
	fmt.Println("Server is Up!")
	// Start the Fiber web server on port 8081
	app.Listen(":8081")
}
