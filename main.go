package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
	//	set content-type header to application header
	c.Header("Content-Type", "application/json")

	//	Display all todos
	err := json.NewEncoder(c.Writer).Encode(todo)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
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

func getTodoByID(id string) (*Todo, error) {
	for key, val := range todo {
		if val.ID == id {
			return &todo[key], nil
		}
	}

	return nil, errors.New("Todo not found")
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

	//listen and serve
	r.Run()
}
