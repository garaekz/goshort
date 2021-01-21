package page

import (
	"context"
	"time"

	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for pages.
type Service interface {
	Get(ctx context.Context, id string) (Link, error)
}

// Link represents the data about an link.
type Link struct {
	Code        string    `json:"code"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new page service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the page with the specified the page ID.
func (s service) Get(ctx context.Context, code string) (Link, error) {
	link, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return Link{}, err
	}

	return Link{Code: link.Code, OriginalURL: link.OriginalURL, CreatedAt: link.CreatedAt}, nil
}
