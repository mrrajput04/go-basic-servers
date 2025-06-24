package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-server/app/handler"
)

func main() {
	fileserver := http.FileServer(http.Dir("../static"))

	http.Handle("/", fileserver)
	http.HandleFunc("/hello", handler.HelloHandler)
	http.HandleFunc("/form", handler.FormHandler)

	fmt.Println("server running on http://localhost:8080")

	if err := http.ListenAndServe("8080", nil); err != nil {
		log.Fatal(err)
	}

}
