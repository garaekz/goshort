package url

// Service has all the repository info
type Service struct {
	Repository Repository
}

// ProvideService provide this service info
func ProvideService(repo Repository) Service {
	return Service{Repository: repo}
}

// FindByCode calls the repository function FindByCode
func (s *Service) FindByCode(code string) (URL, bool) {

	return s.Repository.FindByCode(code)
}

// Save calls the repository function Save
func (s *Service) Save(url URL) URL {
	u, status := s.Repository.FindByOriginalURL(url.OriginalURL)

	if status == false {
		s.Repository.Save(url)
		return url
	}
	return u
}
