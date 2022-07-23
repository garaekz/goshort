package my

import (
	"context"
	"time"

	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/short"
	"github.com/garaekz/goshort/pkg/log"
)

// Service encapsulates usecase logic for shorts.
type Service interface {
	GetMyUser(ctx context.Context) (struct{ User }, error)
	GetMyShorts(ctx context.Context) ([]Short, error)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new short service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Short represents the data about an short.
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Short struct {
	short.ShortResponse
}

// Get returns the short with provided code.
func (s service) GetMyUser(ctx context.Context) (struct{ User }, error) {
	identity := auth.CurrentUser(ctx)
	id := identity.GetID()

	res, err := s.repo.Get(ctx, id)
	user := User{
		ID:        res.ID,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
	}

	if err != nil {
		return struct{ User }{}, err
	}
	return struct{ User }{user}, nil
}

func (s service) GetMyShorts(ctx context.Context) ([]Short, error) {
	identity := auth.CurrentUser(ctx)
	id := identity.GetID()

	res, err := s.repo.GetShortsByOwner(ctx, id)
	if err != nil {
		return []Short{}, err
	}
	return res, nil
}
