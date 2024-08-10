package main

import (
	"fmt"
	"todo_app/internal/db"
)

func main() {
	// Initialize the database connection
	db.InitDB("appuser:123456@tcp(127.0.0.1:3306)/todo_app")
	defer db.CloseDB()

	// Create a new todo item
	id, err := db.CreateTodo("Learn Go", "Read Go documentation")
	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}
	fmt.Printf("New todo created with ID: %d\n", id)

	// Retrieve all todo items
	todos, err := db.GetTodos()
	if err != nil {
		fmt.Println("Error retrieving todos:", err)
		return
	}
	for _, todo := range todos {
		fmt.Printf("Todo ID: %d, Title: %s, Description: %s, Completed: %v\n",
			todo.ID, todo.Title, todo.Description, todo.Completed)
	}
}
