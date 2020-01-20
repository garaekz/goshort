package url

import (
	"fmt"
	"regexp"
	"strings"
)

// Service has all the repository info
type Service struct {
	Repository Repository
}

// ProvideService provide this service info
func ProvideService(repo Repository) Service {
	return Service{Repository: repo}
}

// pruneURL trims URL trailing slash
func (u *URL) pruneURL() {
	if !strings.HasPrefix(u.OriginalURL, "http://") && !strings.HasPrefix(u.OriginalURL, "https://") {
		u.OriginalURL = fmt.Sprintf("http://%s", u.OriginalURL)
	}
	u.OriginalURL = strings.TrimSuffix(u.OriginalURL, "/")
}

// IsURL checks URL validity
func IsURL(str string) bool {
	re := regexp.MustCompile(`^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)?[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`)

	return re.FindStringIndex(str) != nil
}

// setRandomCode sets URL struct Code field
func (u *URL) setRandomCode(s *Service) {
	u.Code = s.Repository.generateUniqueCode()
}

// FindByCode calls the repository function FindByCode
func (s *Service) FindByCode(code string) (URL, bool) {

	return s.Repository.FindByCode(code)
}

// Save calls the repository function Save
func (s *Service) Save(u URL) (URL, bool) {
	u.pruneURL()
	ou, status := s.Repository.FindByOriginalURL(u.OriginalURL)
	fmt.Println(status)
	if !status {
		u.setRandomCode(s)
		fmt.Println(u)
		valid := IsURL(u.OriginalURL)
		if !valid {
			return u, false
		}
		s.Repository.Save(u)
		return u, true
	}
	return ou, true
}
