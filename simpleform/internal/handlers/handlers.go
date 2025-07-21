package handlers

import (
	"context"
	"os"
	"simpleform/internal/model"
	"simpleform/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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
	newUser := model.User{
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

// ReadUsers handles the retrieval of all users from the database
func ReadUsers(c *fiber.Ctx) error {
	// Create a context with a timeout of 7 seconds to ensure the operation does not hang indefinitely
	ctx, cancel := context.WithTimeout(context.TODO(), 7*time.Second)
	defer cancel() // Ensure the context is cancelled to free resources

	// Get the MongoDB database instance from the context
	db := c.Locals("db").(*mongo.Database)

	// Perform a find operation on the user collection to retrieve all user documents
	row, err := db.Collection(os.Getenv("USER_COLLECTION")).Find(ctx, bson.M{})
	if err != nil {
		// Return a bad request error if the find operation fails
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error finding the user. try again" + err.Error(),
		})
	}

	// Declare a slice to hold the user documents
	var users []bson.M

	// Decode all the user documents into the users slice
	if err := row.All(ctx, &users); err != nil {
		// Return a bad request error if decoding the documents fails
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error parsing the user" + err.Error(),
		})
	}

	// Return a success response with the list of all users
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "ALL Users",
		"Data":    users,
	})

}

func ReadOneUser(c *fiber.Ctx) error {
	// Create a context with a timeout of 7 seconds to ensure the operation does not hang indefinitely
	ctx, cancel := context.WithTimeout(context.TODO(), 7*time.Second)
	defer cancel() // Ensure the context is cancelled to free resources

	// Get the user ID from the URL parameters
	id := c.Params("id")

	// Convert the user ID from a hex string to an ObjectID
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Return a bad request error if the ID is not a valid ObjectID
		c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "error validating id, try again " + err.Error(),
		})
	}

	// Declare a variable to hold the user document
	var user model.User
	// Get the MongoDB database instance from the context
	db := c.Locals("db").(*mongo.Database)

	// Find the user document by its ID and decode it into the user variable
	if err := db.Collection(os.Getenv("USER_COLLECTION")).FindOne(ctx, bson.M{"_id": userId}).Decode(&user); err != nil {
		// Return a bad request error if the user is not found or decoding fails
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "error getting user, try again" + err.Error(),
		})
	}

	// Return a success response with the user document
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "scuccessful!",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	// Get the user ID from the URL parameters
	id := c.Params("id")

	// Convert the user ID from a hex string to an ObjectID
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		// Return a bad request error if the ID is not a valid ObjectID
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Error validating ID, try again " + err.Error(),
		})
	}

	form, err := c.MultipartForm()
	if err != nil {
		// Return a bad request error if form data parsing fails
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Error parsing form data: " + err.Error(),
		})
	}

	// Create a map to hold the update data
	updateData := bson.M{}

	// Process the uploaded avatar files
	files := form.File["avatar"]
	var avatarURL string

	for _, fileHead := range files {
		// Open the uploaded file
		file, err := fileHead.Open()
		if err != nil {
			// Return an internal server error if file opening fails
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to open image file: " + err.Error(),
			})
		}
		defer file.Close()

		// Upload the file and get the URL
		avatarURL, err = utils.UTC(file)
		if err != nil {
			// Return an internal server error if file upload fails
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		// Add the avatar URL to the update data
		updateData["avatar"] = avatarURL
	}

	// Get the email and city from the form data and add to update data if not empty
	email := c.FormValue("email")
	if email != "" {
		updateData["email"] = email
	}

	city := c.FormValue("city")
	if email != "" {
		updateData["city"] = city
	}

	// Add the current time to the update data
	updateData["updated_at"] = time.Now()

	// Get the MongoDB database instance from the context
	db := c.Locals("db").(*mongo.Database)
	collection := db.Collection(os.Getenv("USER_COLLECTION"))
	filter := bson.M{"_id": userId}      // Create a filter to find the user by ID
	update := bson.M{"$set": updateData} // Create an update document with the update data

	// Execute the update operation
	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		// Return an internal server error if the update operation fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Error updating user: " + err.Error(),
		})
	}

	// Check if the user was found and updated
	if res.MatchedCount == 0 {
		// Return a not found error if no user was matched
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "User not found",
		})
	}

	// Return a success response if the user was updated successfully
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "User updated successfully",
	})
}

// DeleteUser handles the deletion of a user from the database by their ID
func DeleteUser(c *fiber.Ctx) error {
	// Get the user ID from the URL parameters
	id := c.Params("id")

	// Convert the user ID from a hex string to an ObjectID
	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid user ID",
		})
	}

	// Get the MongoDB database instance from the context
	db := c.Locals("db").(*mongo.Database)

	// Get the user collection from the database
	collection := db.Collection(os.Getenv("USER_COLLECTION"))

	// Create a filter to find the user by ID
	filter := bson.M{"_id": userId}

	// Execute the delete operation
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		// Return an internal server error if the delete operation fails
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Error deleting user: " + err.Error(),
		})
	}
	// Check if the user was found and deleted
	if res.DeletedCount == 0 {
		// Return a not found error if no user was deleted
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Error deleting user: " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "User deleted successfully",
	})
}
