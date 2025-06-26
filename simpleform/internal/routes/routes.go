package routes

import (
	"simpleform/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	route := app.Group("/api/v1/user")
	route.Post("/", handlers.AddUser)
	route.Get("/", handlers.ReadUsers)
}
