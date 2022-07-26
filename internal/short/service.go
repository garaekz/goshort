package short

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/garaekz/goshort/pkg/utils"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Service encapsulates usecase logic for shorts.
type Service interface {
	Get(ctx context.Context, id string) (Short, error)
	GetOwned(ctx context.Context, userID string) ([]Short, error)
	Query(ctx context.Context, offset, limit int) ([]Short, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateShortRequest) (Short, error)
	Update(ctx context.Context, id string, input UpdateShortRequest) (Short, error)
	Delete(ctx context.Context, id string) (Short, error)
	GetCreated(ctx context.Context, id string) (Short, error)
	RegisterVisit(ctx context.Context, code string) (Short, error)
}

// Short represents the data about an short.
type Short struct {
	ShortResponse
}

type ShortResponse struct {
	Code        string    `json:"code"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateShortRequest represents an short creation request.
type CreateShortRequest struct {
	URL string `json:"url"`
	IP  string
}

// Validate validates the CreateShortRequest fields.
func (m CreateShortRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.URL, validation.Required, validation.Length(0, 128)),
		validation.Field(&m.IP, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateShortRequest represents an short update request.
type UpdateShortRequest struct {
	URL string `json:"url"`
}

// Validate validates the CreateShortRequest fields.
func (m UpdateShortRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.URL, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new short service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the short with provided code.
func (s service) Get(ctx context.Context, code string) (Short, error) {
	short, err := s.repo.Get(ctx, code)
	if err != nil {
		return Short{}, err
	}
	return Short{ParseShortResponse(short)}, nil
}

// GetOwned returns owned apiKeys.
func (s service) GetOwned(ctx context.Context, userID string) ([]Short, error) {
	short, err := s.repo.GetOwned(ctx, userID)
	if err != nil {
		return []Short{}, err
	}
	tmpShorts, err := json.Marshal(short)
	if err != nil {
		return []Short{}, err
	}
	var shorts []Short
	err = json.Unmarshal(tmpShorts, &shorts)

	return shorts, nil
}

// GetCreated returns a recently created short with the specified the short ID.
func (s service) GetCreated(ctx context.Context, code string) (Short, error) {
	short, err := s.repo.Get(ctx, code)
	if err != nil {
		return Short{}, err
	}
	return Short{ParseShortResponse(short)}, nil
}

// Create creates a new short.
func (s service) Create(ctx context.Context, req CreateShortRequest) (Short, error) {
	if err := req.Validate(); err != nil {
		return Short{}, err
	}

	identity := auth.CurrentUser(ctx)
	userID := identity.GetID()

	// Making sure the URL is in a good format
	URL, err := utils.FormatURL(req.URL)
	if err != nil {
		return Short{}, errors.New("URL is not in a valid format")
	}

	// Finding out if this URL has been already saved in the DB as an OriginalURL
	short, err := s.repo.GetByOriginalURL(ctx, URL, userID)

	if err != nil {
		if err != sql.ErrNoRows {
			return Short{}, err
		}
	}
	// If it's already on the DB we just return that
	if (Short{ParseShortResponse(short)}) != (Short{}) {
		return Short{ParseShortResponse(short)}, nil
	}

	code, err := s.repo.GenerateUniqueCode(ctx)
	if err != nil {
		return Short{}, err
	}

	now := time.Now()
	err = s.repo.Create(ctx, entity.Short{
		Code:        code,
		OriginalURL: URL,
		Visits:      0,
		UserID:      userID,
		CreatorIP:   req.IP,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	if err != nil {
		return Short{}, err
	}

	return s.GetCreated(ctx, code)
}

// Update updates the short with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateShortRequest) (Short, error) {
	if err := req.Validate(); err != nil {
		return Short{}, err
	}

	short, err := s.repo.Get(ctx, id)
	if err != nil {
		return Short{}, err
	}

	URL, err := utils.FormatURL(req.URL)
	if err != nil {
		return Short{}, errors.New("URL is not in a valid format")
	}
	short.OriginalURL = URL
	short.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, short); err != nil {
		return Short{}, err
	}
	return Short{ParseShortResponse(short)}, nil
}

// Delete deletes the short with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Short, error) {
	short, err := s.Get(ctx, id)
	if err != nil {
		return Short{}, err
	}

	if err = s.repo.Delete(ctx, id); err != nil {
		return Short{}, err
	}

	return short, nil
}

// Count returns the number of shorts.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the shorts with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Short, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Short{}
	for _, item := range items {
		result = append(result, Short{ParseShortResponse(item)})
	}
	return result, nil
}

// RegisterVisit updates the short visit count with the specified code.
func (s service) RegisterVisit(ctx context.Context, code string) (Short, error) {
	short, err := s.repo.Get(ctx, code)
	if err != nil {
		return Short{}, err
	}
	short.Visits++

	if err := s.repo.Update(ctx, short); err != nil {
		return Short{}, err
	}
	return Short{ParseShortResponse(short)}, nil
}

// ParseShortResponse parses a Short entity into a secure response.
func ParseShortResponse(original entity.Short) ShortResponse {
	return ShortResponse{
		Code:        original.Code,
		OriginalURL: original.OriginalURL,
		CreatedAt:   original.CreatedAt,
		UpdatedAt:   original.UpdatedAt,
	}
}
