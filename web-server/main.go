package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// data definitions
type person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// people slice to seed record person data.
var people = []person{
	{
		ID:   "1",
		Name: "ABC",
	},
	{
		ID:   "2",
		Name: "DEF",
	},
	{
		ID:   "3",
		Name: "GHI",
	},
}

func main() {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8888 (for windows "localhost:8888")
	// -- snip --
	r.GET("/people", getPeople)
	r.POST("/people", postPeople)
	r.GET("/people/:id", getPersonByID)
	// -- snip --
}

// respond with the entire people struct as JSON
func getPeople(context *gin.Context) {
	// IndentedJSON makes it look better
	context.IndentedJSON(http.StatusOK, people)
}

// add a person to people from JSON received in the request body
func postPeople(context *gin.Context) {
	var newPerson person
	// BindJSON to bind the received JSON to newPerson
	if err := context.BindJSON(&newPerson); err != nil {
		// log the error, respond and return
		fmt.Println(err)
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	// append the new person to people
	people = append(people, newPerson)
	// respond as IndentedJSON
	context.IndentedJSON(http.StatusCreated, newPerson)
}

// locate the person whose ID value matches the id sent
// then return that person as a response
func getPersonByID(context *gin.Context) {
	// get the id from request params
	var id string = context.Param("id")

	// Linear Search through people
	for _, p := range people {
		// respond and return if ID matched
		if p.ID == id {
			context.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	// respond 404
	context.IndentedJSON(
		http.StatusNotFound,
		gin.H{
			"message": "person not found",
		},
	)
}
