package db

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
)

// DB is a global variable to hold the database connection pool
var DB *sql.DB

// InitDB initializes the database connection
func InitDB(dataSourceName string) {
    var err error
    DB, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err)
    }

    // Check if the connection is alive
    if err := DB.Ping(); err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    log.Println("Database connection successfully established")
}

// CloseDB closes the database connection
func CloseDB() {
    if err := DB.Close(); err != nil {
        log.Fatalf("Error closing database connection: %v", err)
    }
}

// CreateTodo inserts a new todo item into the database
func CreateTodo(title, description string) (int64, error) {
    result, err := DB.Exec("INSERT INTO todos (title, description) VALUES (?, ?)", title, description)
    if err != nil {
        return 0, err
    }

    // Get the ID of the newly inserted row
    lastInsertID, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return lastInsertID, nil
}

// GetTodos retrieves all todo items from the database
func GetTodos() ([]Todo, error) {
    rows, err := DB.Query("SELECT id, title, description, completed FROM todos")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var todos []Todo
    for rows.Next() {
        var todo Todo
        if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed); err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return todos, nil
}

// Todo represents the structure of a todo item
type Todo struct {
    ID          int64
    Title       string
    Description string
    Completed   bool
}

//User info 
