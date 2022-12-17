package models

import (
	"time"
)

type (
	// Cekes
	Cake struct {
		ID          int       `json:"id"`
		Title       string    `json:"title"`
		Description string    `name:"description"`
		Rating      int       `json:"rating"`
		Images      string    `name:"images"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
