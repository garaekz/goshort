package entity

import "time"

// ApiKey represents a user record.
type ApiKey struct {
	Key       string     `json:"key" db:"key"`
	UserID    string     `json:"user_id" db:"user_id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

// TableName represents the table name
func (model ApiKey) TableName() string {
	return "api_keys"
}

// GetKey returns the user ID.
func (model ApiKey) GetKey() string {
	return model.Key
}
