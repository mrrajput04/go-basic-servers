package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-server/app/handler"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", handler.HelloHandler)
	http.HandleFunc("/form", handler.FormHandler)

	fmt.Printf("port running on http://localhost:8081/\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
