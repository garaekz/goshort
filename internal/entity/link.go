package entity

import (
	"time"
)

// Link represents a link record.
type Link struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	OriginalURL string    `json:"original_url"`
	UserID      *string   `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName represents the table name
func (u Link) TableName() string {
	return "urls"
}
