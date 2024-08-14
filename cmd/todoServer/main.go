package main

import (
	"log"

	"todoServer/internal/config"
	"todoServer/internal/db"
	"todoServer/internal/gee"
	"todoServer/internal/todo_handler"
)

func main() {
	// 读取配置
	config.LoadConfig()

	// 初始化数据库
	err := db.InitDB(config.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 运行数据库迁移
	err = db.RunMigrations(db.GetDB())
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// // CORS 配置
	// corsHandler := handlers.CORS(
	// 	handlers.AllowedOrigins([]string{"http://localhost:3000"}), // 允许的源
	// 	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	// 	handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	// )(r)
	// http.Handle("/", corsHandler)

	r := gee.New()
	//todo路由
	r.GET("/getTodos", todo_handler.GetTodos)
	r.POST("/postTodo", todo_handler.PostTodo)

	// user路由
	r.POST("/registerUser", todo_handler.RegisterUser)
	r.POST("/loginUser", todo_handler.LoginUser)

	// 启动 HTTP 服务器
	r.Run(":80")
	log.Println("Starting server on :80")
}
