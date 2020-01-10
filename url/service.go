package url

import (
	"fmt"
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
func (url *URL) pruneURL() {
	if !strings.HasPrefix(url.OriginalURL, "http://") && !strings.HasPrefix(url.OriginalURL, "https://") {
		url.OriginalURL = fmt.Sprintf("http://%s", url.OriginalURL)
	}
	url.OriginalURL = strings.TrimSuffix(url.OriginalURL, "/")
}

// setRandomCode sets URL struct Code field
func (url *URL) setRandomCode(s *Service) {
	url.Code = s.Repository.generateUniqueCode()
}

// FindByCode calls the repository function FindByCode
func (s *Service) FindByCode(code string) (URL, bool) {

	return s.Repository.FindByCode(code)
}

// Save calls the repository function Save
func (s *Service) Save(url URL) URL {
	url.pruneURL()
	u, status := s.Repository.FindByOriginalURL(url.OriginalURL)

	if status == false {
		url.setRandomCode(s)
		s.Repository.Save(url)
		return url
	}
	return u
}
