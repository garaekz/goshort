package api_key

import (
	"context"
	"time"

	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/log"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Service encapsulates usecase logic for api_keys.
type Service interface {
	Get(ctx context.Context, id string) (ApiKey, error)
	Query(ctx context.Context, offset, limit int) ([]ApiKey, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateApiKeyRequest) (ApiKey, error)
	Update(ctx context.Context, id string, input UpdateApiKeyRequest) (ApiKey, error)
	Delete(ctx context.Context, id string) (ApiKey, error)
}

// ApiKey represents the data about an api_key.
type ApiKey struct {
	entity.ApiKey
}

// CreateApiKeyRequest represents an api_key creation request.
type CreateApiKeyRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate validates the CreateApiKeyRequest fields.
func (m CreateApiKeyRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Required),
	)
}

// UpdateApiKeyRequest represents an api_key update request.
type UpdateApiKeyRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate validates the CreateApiKeyRequest fields.
func (m UpdateApiKeyRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Email, validation.Required, is.Email),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new api_key service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the api_key with the specified the api_key ID.
func (s service) Get(ctx context.Context, id string) (ApiKey, error) {
	api_key, err := s.repo.Get(ctx, id)
	if err != nil {
		return ApiKey{}, err
	}
	return ApiKey{api_key}, nil
}

// Create creates a new api_key.
func (s service) Create(ctx context.Context, req CreateApiKeyRequest) (ApiKey, error) {

	if err := req.Validate(); err != nil {
		return ApiKey{}, err
	}

	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.ApiKey{
		Key:       id,
		CreatedAt: now,
	})
	if err != nil {
		return ApiKey{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the api_key with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateApiKeyRequest) (ApiKey, error) {
	if err := req.Validate(); err != nil {
		return ApiKey{}, err
	}

	api_key, err := s.Get(ctx, id)
	if err != nil {
		return api_key, err
	}

	if err := s.repo.Update(ctx, api_key.ApiKey); err != nil {
		return api_key, err
	}
	return api_key, nil
}

// Delete deletes the api_key with the specified ID.
func (s service) Delete(ctx context.Context, id string) (ApiKey, error) {
	api_key, err := s.Get(ctx, id)
	if err != nil {
		return ApiKey{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return ApiKey{}, err
	}
	return api_key, nil
}

// Count returns the number of api_keys.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the api_keys with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]ApiKey, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []ApiKey{}
	for _, item := range items {
		result = append(result, ApiKey{item})
	}
	return result, nil
}
