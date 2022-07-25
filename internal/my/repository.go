package my

import (
	"context"

	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/dbcontext"
	"github.com/garaekz/goshort/pkg/log"
)

// Repository encapsulates the logic to access shorts from the data source.
type Repository interface {
	// Get returns the short with the specified short ID.
	Get(ctx context.Context, id string) (entity.User, error)
}

// repository persists shorts in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new short repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the user with the specified ID from the database.
func (r repository) Get(ctx context.Context, userID string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().Model(userID, &user)
	return user, err
}
