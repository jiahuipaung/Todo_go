package models

import (
	"time"
)

type Todo struct {
	ItemId          uint      `json:"itemId"`
	UID             uint      `json:"uid"`
	ItemName        string    `json:"itemName"`
	Description     string    `json:"description"`
	Completed       bool      `json:"completed"`
	CreatedAt       time.Time `json:"createdAt"`
	ItemDeadline    string    `json:"itemDeadline"`
	NeedCheckInDays uint      `json:"needCheckInDays"`
}
