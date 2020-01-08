package url

import (
	"github.com/jinzhu/gorm"
)

// Repository passes current database connection
type Repository struct {
	DB *gorm.DB
}

// ProvideRepository returns current repository
func ProvideRepository(DB *gorm.DB) Repository {
	return Repository{DB: DB}
}

// FindByCode finds the given Code url
func (repo *Repository) FindByCode(code string) (URL, bool) {
	var url URL

	if repo.DB.Where("code = ?", code).First(&url).RecordNotFound() {
		return url, false
	}

	return url, true
}

// FindByOriginalURL finds the given OriginalURL url
func (repo *Repository) FindByOriginalURL(originalURL string) (URL, bool) {
	var url URL

	if repo.DB.Where("original_url = ?", originalURL).First(&url).RecordNotFound() {
		return url, false
	}

	return url, true
}

// Save saves the url model to database and publish it to redis, redis part is not ready yet
func (repo *Repository) Save(url URL) URL {
	repo.DB.Save(&url)

	return url
}
