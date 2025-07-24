package main

import (
	"comic-store-apis/config"
	"comic-store-apis/handlers"
	"net/http"
)

func main() {
	// Connect to the database using the configuration setup
	config.ConnectDB()

	// Register the AppHandler function to handle requests at the root path "/"
	http.HandleFunc("/", handlers.AppHandler)

	// Start the HTTP server and listen on port 8000
	http.ListenAndServe(":8888", nil)

}
