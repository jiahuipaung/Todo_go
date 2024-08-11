package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 假设这里有身份验证逻辑
		next.ServeHTTP(w, r)
	})
}
