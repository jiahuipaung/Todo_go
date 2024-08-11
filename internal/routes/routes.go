package routes

import (
	"net/http"
	"todo_app/internal/middleware"
	"todo_app/internal/myhandlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Todo 路由
	r.HandleFunc("/getTodos", myhandlers.GetTodos).Methods("GET")
	r.HandleFunc("/post_todo", handlers.PostTodos).Methods("POST")

	// CORS 配置
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // 允许的源
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	http.Handle("/", corsHandler)
	// 用户路由（如果有用户功能）
	r.HandleFunc("/users", myhandlers.GetUsers).Methods("GET")

	// 使用身份验证中间件
	r.Use(middleware.AuthMiddleware)

	return r
}
