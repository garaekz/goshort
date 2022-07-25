package entity

import (
	"time"
)

// APIKey represents a key record.
type APIKey struct {
	Key       string    `json:"key" db:"pk,key"`
	UserID    string    `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	IsActive  bool      `json:"is_active" db:"-"`
}

// TableName represents the table name
func (k APIKey) TableName() string {
	return "keys"
}
