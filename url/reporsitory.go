package url

import (
	"crypto/rand"

	"github.com/jinzhu/gorm"
)

// StdLen sets the initial length of the generated code
const (
	StdLen = 4
)

// StdChars contains the valid characters to use in a code
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

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

	if repo.DB.Where("original_url ilike ?", originalURL).First(&url).RecordNotFound() {
		return url, false
	}

	return url, true
}

// Save saves the url model to database and publish it to redis, redis part is not ready yet
func (repo *Repository) Save(url URL) URL {
	repo.DB.Save(&url)

	return url
}

func (repo *Repository) generateUniqueCode() string {
	n := StdLen
	i := 0

	for {
		code := repo.RandomCode(n)
		_, status := repo.FindByCode(code)
		if !status {
			return code
		}
		if n%10 == 0 {
			n++
		}
		i++
	}
}

// RandomCode returns a new random string of the provided length.
func (repo *Repository) RandomCode(length int) string {
	return repo.RandomCodeChars(length, StdChars)
}

// RandomCodeChars returns a new random string of the provided length using the byte slice.
func (repo *Repository) RandomCodeChars(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("uniuri: wrong charset length for NewLenChars")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4))
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("uniuri: error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}
