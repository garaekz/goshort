package link

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/log"
)

// Service encapsulates usecase logic for links.
type Service interface {
	Get(ctx context.Context, code string) (Link, error)
	Query(ctx context.Context, offset, limit int) ([]Link, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateLinkRequest) (Link, error)
	Update(ctx context.Context, id string, input UpdateLinkRequest) (Link, error)
	Delete(ctx context.Context, id string) (Link, error)
}

// Link represents the data about an link.
type Link struct {
	Code        string    `json:"code"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateLinkRequest represents an link creation request.
type CreateLinkRequest struct {
	URL string `json:"url"`
}

// Validate validates the CreateLinkRequest fields.
func (m CreateLinkRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.URL, validation.Required),
	)
}

// UpdateLinkRequest represents an link update request.
type UpdateLinkRequest struct {
	Code        string    `json:"code"`
	OriginalURL string    `json:"original_url"`
	UserID      *string   `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Validate validates the CreateLinkRequest fields.
func (m UpdateLinkRequest) Validate() error {
	return validation.ValidateStruct(&m) //validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),

}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new link service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the link with the specified the link ID.
func (s service) Get(ctx context.Context, id string) (Link, error) {
	link, err := s.repo.Get(ctx, id)
	if err != nil {
		return Link{}, err
	}
	return Link{Code: link.Code, OriginalURL: link.OriginalURL, CreatedAt: link.CreatedAt}, nil
}

// Create creates a new link.
func (s service) Create(ctx context.Context, req CreateLinkRequest) (Link, error) {
	var userID *string
	if err := req.Validate(); err != nil {
		return Link{}, err
	}

	identity := auth.CurrentUser(ctx)
	if identity != nil {
		identityID := identity.GetID()
		userID = &identityID
	}

	// Making sure the URL is in a good format
	URL, err := formatURL(req.URL)
	if err != nil {
		return Link{}, err
	}

	// Finding out if this URL has been already saved in the DB as an OriginalURL
	link, err := s.repo.GetByOriginalURL(ctx, URL)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return Link{}, err
		}
	}

	// If it's already on the DB we just return that
	if link != nil {
		return Link{Code: link.Code, OriginalURL: link.OriginalURL, CreatedAt: link.CreatedAt}, nil
	}

	code, err := s.repo.GenerateUniqueCode(ctx)
	if err != nil {
		return Link{}, err
	}

	id := entity.GenerateID()
	now := time.Now()
	err = s.repo.Create(ctx, entity.Link{
		ID:          id,
		Code:        code,
		UserID:      userID,
		OriginalURL: URL,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		return Link{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the link with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateLinkRequest) (Link, error) {
	if err := req.Validate(); err != nil {
		return Link{}, err
	}

	return Link{}, nil
	/* 	link, err := s.Get(ctx, id)
	   	if err != nil {
	   		return link, err
	   	}
	   	link.Name = req.Name
	   	link.UpdatedAt = time.Now()

	   	if err := s.repo.Update(ctx, link.Link); err != nil {
	   		return link, err
	   	}
	   	return link, nil */
}

// Delete deletes the link with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Link, error) {
	link, err := s.Get(ctx, id)
	if err != nil {
		return Link{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Link{}, err
	}
	return link, nil
}

// Count returns the number of links.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the links with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Link, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Link{}
	for _, item := range items {
		result = append(result, Link{Code: item.Code, OriginalURL: item.OriginalURL, CreatedAt: item.CreatedAt})
	}
	return result, nil
}
