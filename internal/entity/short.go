package entity

import (
	"time"
)

// Short represents an album record.
type Short struct {
	Code        string     `db:"pk,code"`
	OriginalURL string     `db:"original_url"`
	Visits      int        `db:"visits"`
	UserID      string     `db:"user_id"`
	CreatorIP   string     `db:"creator_ip"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

// TableName represents the table name
func (s Short) TableName() string {
	return "shorts"
}
