package myhandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"todo_app/internal/db"
	"todo_app/internal/models"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	// 从查询参数中获取 uid
	uid := r.URL.Query().Get("uid")

	// // 如果没有提供 uid，返回错误
	// if uid ==  {
	// 	http.Error(w, "Missing uid parameter", http.StatusBadRequest)
	// 	return
	// }

	fmt.Printf("uid=%s \n", uid)

	// 连接数据库
	dbConn := db.GetDB()

	// 查询数据库
	query := "SELECT itemId, uid, itemName FROM todos WHERE uid = ?"
	rows, err := dbConn.Query(query, uid)
	if err != nil {
		log.Printf("Error querying todos: %v", err)
		http.Error(w, "Error querying todos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	todos := []models.Todo{}
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ItemId, &todo.UID, &todo.ItemName); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	// 检查查询是否遇到错误
	if err := rows.Err(); err != nil {
		log.Printf("Error in row iteration: %v", err)
		http.Error(w, "Error processing query results", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func PostTodos(w http.ResponseWriter, r *http.Request) {
	//解析todo数据
	var todo models.Todo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// 设置默认值
	if todo.CreatedAt.IsZero() {
		todo.CreatedAt = time.Now()
	}

	

	// 连接数据库
	dbConn := db.GetDB()

	// 插入 TODO 到数据库
	query := `INSERT INTO todos (itemId, uid, itemName, description, completed, created_at, itemDeadline, needCheckInDays) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = dbConn.Exec(query, todo.ItemId, todo.UID, todo.ItemName, todo.Description, todo.Completed, todo.CreatedAt, todo.ItemDeadline, todo.NeedCheckInDays)
	if err != nil {
		http.Error(w, "Failed to insert TODO", http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
