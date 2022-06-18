package entity

import "time"

type Brand struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Logo string `json:"logo"`
	Level string `json:"level"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}