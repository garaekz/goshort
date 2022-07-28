package entity

import "time"

// EmailVerification represents a user.
type EmailVerification struct {
	UserID    string    `json:"user_id" db:"pk,user_id"`
	Token     string    `json:"token" db:"pk,token"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

// TableName represents the table name
func (v EmailVerification) TableName() string {
	return "email_verifications"
}
