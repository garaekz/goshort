package url

import (
	"github.com/jinzhu/gorm"
)

// URL represents a database table
type URL struct {
	gorm.Model
	Code        string `gorm:"unique;not null" json:"code"`
	OriginalURL string `gorm:"not null" json:"original_url"`
}
