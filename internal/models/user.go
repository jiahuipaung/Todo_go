package models

import (
	"time"
)

type User struct {
	UID          uint      `json:"uid"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}
