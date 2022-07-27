package entity

import "time"

// User represents a user.
type User struct {
	ID            string    `json:"id" db:"pk,id"`
	Email         string    `json:"email" db:"email"`
	Password      string    `json:"-" db:"password"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	EmailVerified bool      `json:"email_verified" db:"email_verified"`
	IsActive      bool      `json:"is_active" db:"is_active"`
}

// TableName represents the table name
func (s User) TableName() string {
	return "users"
}

// GetID returns the user ID.
func (u User) GetID() string {
	return u.ID
}

// GetEmail returns the user ID.
func (u User) GetEmail() string {
	return u.Email
}
