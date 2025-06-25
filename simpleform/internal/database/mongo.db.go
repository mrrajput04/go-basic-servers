package database

import (
	"log"

	"github.com/joho/godotenv"
)

func DBClient() *mongo.database {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file, check if it exists.")
	}

	log.Println("Connecting to mongodb....")
	serverAPIOptions := options.ServerAPI(options.serverAPIVersion1)
	clientOptions := options.Client()
}
