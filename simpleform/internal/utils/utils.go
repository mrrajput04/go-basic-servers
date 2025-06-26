package utils

import (
	"context"
	"errors"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

// CorsMiddleware enables Cross-Origin Resource Sharing (CORS) for the application
func CorsMiddleware(app *fiber.App) {
	// Use the fiber.Cors middleware to allow cross-origin requests
	app.Use(cors.New())
	// You can optionally add additional configuration options to the cors.New() function
	// to customize CORS behavior for your application.
}

// DbMiddleware injects the database connection into the request context
func DbMiddleware(db *mongo.Database) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Store the database connection in the context's "db" key
		c.Locals("db", db)
		// Call the next middleware in the chain
		return c.Next()
	}
}

// UTC uploads an image file to Cloudinary and returns the secure URL (likely a misnamed function)
func UTC(file multipart.File) (string, error) {
	urlCloudinary := os.Getenv("CLOUDINARY_URL")
	cloudService, err := cloudinary.NewFromURL(urlCloudinary)
	if err != nil {
		return "", errors.New("failed to create Cloudinary service" + err.Error())
	}

	ctx := context.Background()
	resp, err := cloudService.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		return "", errors.New("failed to uplaod image to cloudinary" + err.Error())
	}

	return resp.SecureURL, nil
}
