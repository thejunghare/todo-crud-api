package main

import "fmt"

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
