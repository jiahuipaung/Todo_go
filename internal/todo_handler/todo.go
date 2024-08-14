package todo_handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"todoServer/internal/db"
	"todoServer/internal/gee"
	"todoServer/internal/models"
)

func GetTodos(c *gee.Context) {
	// 从查询参数中获取 uid
	uid := c.Query("uid")

	// 如果没有提供 uid，返回错误
	if uid == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	// 连接数据库
	dbConn := db.GetDB()

	// 查询数据库
	query := "SELECT itemId, uid, itemName FROM todos WHERE uid = ?"
	rows, err := dbConn.Query(query, uid)
	if err != nil {
		log.Printf("Error querying todos: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	todos := []models.Todo{}
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ItemId, &todo.UID, &todo.ItemName); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}

	// 检查查询是否遇到错误
	if err := rows.Err(); err != nil {
		log.Printf("Error in row iteration: %v", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, todos)
}

func PostTodo(c *gee.Context) {
	//解析todo数据
	var todo models.Todo
	c.PostForm()

	// 设置默认值
	if todo.CreatedAt.IsZero() {
		todo.CreatedAt = time.Now()
	}

	// 连接数据库
	dbConn := db.GetDB()

	// 定义时间字符串和对应的布局
	layout := time.RFC3339 // 标准 ISO 8601 格式布局

	// 解析时间字符串
	itemDeadlineTimeStamp, err := time.Parse(layout, todo.ItemDeadline)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	// 插入 TODO 到数据库
	query := `INSERT INTO todos (uid, itemName, description, created_at, itemDeadline, needCheckInDays) 
              VALUES (?, ?, ?, ?, ?, ?)`
	_, err = dbConn.Exec(query, todo.UID, todo.ItemName, todo.Description, todo.CreatedAt, itemDeadlineTimeStamp, todo.NeedCheckInDays)
	if err != nil {
		http.Error(w, "Failed to insert Todo", http.StatusInternalServerError)
		return
	}

	// 返回成功响应
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
