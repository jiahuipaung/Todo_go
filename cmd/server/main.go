package main

import (
	"log"
	"net/http"
	"todo_app/internal/config"
	"todo_app/internal/db"
	"todo_app/internal/routes"
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

	// 设置路由
	r := routes.SetupRoutes()

	// 启动 HTTP 服务器
	log.Println("Starting server on :80")
	err = http.ListenAndServe(":80", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
