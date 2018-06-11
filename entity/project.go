package entity

import (
	"time"
)

// Project entity
type Project struct {
	ID          `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}
