package handlers

import (
	"os"
	"simpleform/internal/models"
	"simpleform/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddUser handles the creation of a new user
func AddUser(c *fiber.Ctx) error {
	// Parse the multipart form data
	form, err := c.MultipartForm()
	if err != nil {
		// Return a bad request error if form parsing fails
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to image file: " + err.Error(),
		})
	}

	// Get the uploaded avatar files from the form data
	files := form.File["avatar"]
	var avatarURL string

	// Process each uploaded file
	for _, fileHead := range files {
		file, err := fileHead.Open()
		if err != nil {
			// Return an internal server error if the file cannot be opened
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "failed to open image file" + err.Error(),
			})
		}
		defer file.Close() // Ensure the file is closed after processing

		// Upload the file to the storage service and get the URL
		avatarURL, err = utils.UTC(file)
		if err != nil {
			// Return an internal server error if the file upload fails
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
	}

	// Get the email and city values from the form data
	email := c.FormValue("email")
	city := c.FormValue("city")
	now := time.Now()

	// Create a new user object with the provided data
	newUser := models.User{
		ID:           primitive.NewObjectID(), // Generate a new MongoDB ObjectID
		EmailAddress: email,                   // Set the email address
		City:         city,                    // Set the city
		Avatar:       avatarURL,               // Set the avatar URL
		CreatedAt:    now,                     // Set the creation timestamp
		UpdatedAt:    now,                     // Set the update timestamp
	}

	// Get the MongoDB database instance from the context
	db := c.Locals("db").(*mongo.Database)

	// Insert the new user into the collection
	res, err := db.Collection(os.Getenv("USER_COLLECTION")).InsertOne(c.Context(), newUser)
	if err != nil {
		// Return an internal server error if the insert operation fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	// Return a success response with the newly created user ID
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": "welcome have a look around",
		"user":    res.InsertedID,
	})

}
