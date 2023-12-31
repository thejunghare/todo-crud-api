package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var todo []Todo

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Todo crud api"})
}

// Display all todo
func getTodos(c *gin.Context) {

	/*
		//	set content-type header to application header
		c.Header("Content-Type", "application/json")

		//	Display all todos
		err := json.NewEncoder(c.Writer).Encode(todo)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	*/

	c.IndentedJSON(http.StatusOK, todo)
}

// Display todo by id
func getTodo(c *gin.Context) {
	// Display by id
	id := c.Param("id")
	todo, err := getTodoByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo does not exist"})
	}

	c.IndentedJSON(http.StatusOK, todo)
}

// Helper function
func getTodoByID(id string) (*Todo, error) {
	for key, val := range todo {
		if val.ID == id {
			return &todo[key], nil
		}
	}

	return nil, errors.New("Todo not found")
}

// Create todo
func createTodo(c *gin.Context) {
	var create Todo

	// Parse the struct
	err := json.NewDecoder(c.Request.Body).Decode(&create)
	if err != nil {
		log.Fatal(err)
	}

	// add todo
	todo = append(todo, create)
	c.IndentedJSON(http.StatusCreated, todo)
}

// update by id
func updateTodo(c *gin.Context) {
	var update Todo
	id := c.Param("id")

	// Find the index of the todo with the specified ID
	index := -1
	for i, t := range todo {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo does not exist"})
		return // Exit the function early if the todo is not found
	}

	// Parse the request body to get the updated todo data
	if err := c.ShouldBindJSON(&update); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return // Exit the function early if there's an error decoding the request body
	}

	// Update the todo data at the found index
	todo[index].Title = update.Title

	c.IndentedJSON(http.StatusOK, todo[index]) // Respond with the updated todo
}

// Delete by id
func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	_, err := getTodoByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo does not exist"})
	}

	for key, val := range todo {
		if val.ID == id {
			todo = append(todo[:key], todo[key+1:]...)
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}

}

func main() {
	fmt.Println("todo crud api")

	// start gin server
	r := gin.Default()

	//	db
	todo = append(todo, Todo{ID: "1", Title: "Create todo crud api"})
	todo = append(todo, Todo{ID: "2", Title: "Create todo crud api with gorm"})

	//	Home Handler
	r.GET("/", HomeHandler)
	//	Get Todo's
	r.GET("/todos", getTodos)
	// Get Todo by ID
	r.GET("/todo/:id", getTodo)
	// Create Todo
	r.PUT("/create", createTodo)
	// Update Todo
	r.PUT("/update/:id", updateTodo)
	// Delete Todo
	r.DELETE("/delete/:id", deleteTodo)

	//listen and serve
	r.Run()
}
