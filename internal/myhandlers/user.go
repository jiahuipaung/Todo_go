package myhandlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"todo_app/internal/db"
	"todo_app/internal/models"
	"todo_app/internal/utils"
)

// func setLoginCookie(w http.ResponseWriter, uid uint) {
// 	// 创建一个包含用户 ID 的 Cookie
// 	cookie := &http.Cookie{
// 		Name:     "uid",
// 		Value:    strconv.FormatUint(uint64(uid), 10), // 将 uid 转换为字符串
// 		Path:     "/",
// 		HttpOnly: true,                           // 防止 JavaScript 访问
// 		Secure:   true,                           // 仅在 HTTPS 下发送
// 		Expires:  time.Now().Add(24 * time.Hour), // 设置过期时间
// 	}

// 	// 设置 Cookie 到响应中
// 	http.SetCookie(w, cookie)
// }

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dbConn := db.GetDB()

	//检查email是否已被注册
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	err := dbConn.QueryRow(query, user.Email).Scan(&count)
	if err != nil {
		return
	}

	// 设置creatAt默认值
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}

	// Insert user into the database
	query = "INSERT INTO users (email, password, created_at) VALUES (?, ?, ?)"
	_, err = dbConn.Exec(query, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	rsp := utils.UtilRsp{Code: "SUCCESS", Message: "register success"}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rsp)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dbConn := db.GetDB()

	//检查email是否已被注册
	var existingUser models.User
	query := "SELECT email, password FROM users WHERE email = ?"
	err := dbConn.QueryRow(query, user.Email).Scan(&existingUser.Email, &existingUser.Password)
	if err != nil {
		rsp := utils.UtilRsp{Code: "FAILED", Message: "Not found the email"}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rsp)
		return
	}

	//检查password
	if user.Password != existingUser.Password {
		rsp := utils.UtilRsp{Code: "FAILED", Message: "password is wrong"}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(rsp)
		return
	}

	rsp := utils.UtilRsp{Code: "SUCCESS", Message: "login success"}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rsp)
}
