package models

import (
	"time"
)

type User struct {
	UID       uint      `json:"uid"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
