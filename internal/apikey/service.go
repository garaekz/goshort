package apikey

import (
	"context"
	"database/sql"
	"time"

	"github.com/garaekz/goshort/internal/entity"
	customErrors "github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/garaekz/goshort/pkg/utils"
)

// Service encapsulates usecase logic for shorts.
type Service interface {
	Get(ctx context.Context, id string) (APIKey, error)
	GetOwned(ctx context.Context, userID string) ([]APIKey, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, UserID string) (APIKey, error)
	Delete(ctx context.Context, id string) (APIKey, error)
	GenerateUniqueKey(ctx context.Context) (string, error)
}

// APIKey represents the data about an API Key.
type APIKey struct {
	Response
}

// Response represents the returned API Key.
type Response struct {
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateAPIKeyRequest represents an API Key creation request.
type CreateAPIKeyRequest struct {
	Key    string `db:"key"`
	UserID string `db:"user_id"`
}

type service struct {
	repo       Repository
	logger     log.Logger
	MaxAPIKeys int
}

// NewService creates a new short service.
func NewService(repo Repository, logger log.Logger, MaxAPIKeys int) Service {
	return service{repo, logger, MaxAPIKeys}
}

// Get returns the short with provided key.
func (s service) Get(ctx context.Context, key string) (APIKey, error) {
	apiKey, err := s.repo.Get(ctx, key)
	if err != nil {
		return APIKey{}, err
	}
	return ParseResponse(apiKey), nil
}

// GetOwned returns owned apiKeys.
func (s service) GetOwned(ctx context.Context, userID string) ([]APIKey, error) {
	apiKeys, err := s.repo.GetOwned(ctx, userID)
	if err != nil {
		return []APIKey{}, err
	}
	return apiKeys, nil
}

// Create creates a new apiKey.
func (s service) Create(ctx context.Context, userID string) (APIKey, error) {
	count, err := s.repo.CountByOwner(ctx, userID)
	if err != nil {
		return APIKey{}, err
	}

	if count >= s.MaxAPIKeys {
		return APIKey{}, customErrors.MaxAPIKeys("You have reached the maximum number of API Keys allowed.")
	}

	key, err := s.GenerateUniqueKey(ctx)
	if err != nil {
		return APIKey{}, err
	}

	apiKey := entity.APIKey{
		Key:       key,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err = s.repo.Create(ctx, apiKey); err != nil {
		return APIKey{}, err
	}

	return s.Get(ctx, key)
}

// Delete deletes the short with the specified ID.
func (s service) Delete(ctx context.Context, key string) (APIKey, error) {
	apiKey, err := s.Get(ctx, key)
	if err != nil {
		return APIKey{}, err
	}

	if err = s.repo.Delete(ctx, key); err != nil {
		return APIKey{}, err
	}

	return apiKey, nil
}

// Count returns the number of apiKeys.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// GenerateUniqueKey generates a unique string.
func (s service) GenerateUniqueKey(ctx context.Context) (string, error) {
	n := 64
	i := 0

	for {
		apiKey, err := utils.RandomString(n, apiKeyChars)
		if err != nil {
			return "", err
		}

		_, err = s.repo.Get(ctx, apiKey)
		if err != nil {
			if err == sql.ErrNoRows {
				return apiKey, nil
			}
			return "", err
		}

		if i%10 == 0 {
			n++
		}
		i++
	}
}

// ParseResponse parses a Short entity into a secure response.
func ParseResponse(original entity.APIKey) APIKey {
	return APIKey{
		Response{
			Key:       original.Key,
			CreatedAt: original.CreatedAt,
		},
	}
}
