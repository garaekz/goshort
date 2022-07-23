package my

import (
	"context"

	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/dbcontext"
	"github.com/garaekz/goshort/pkg/log"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access shorts from the data source.
type Repository interface {
	// Get returns the short with the specified short ID.
	Get(ctx context.Context, id string) (entity.User, error)
	// GetShortsByOwner returns the shorts of the specified user.
	GetShortsByOwner(ctx context.Context, id string) ([]Short, error)
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
func (r repository) Get(ctx context.Context, id string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().Model(id, &user)
	return user, err
}

func (r repository) GetShortsByOwner(ctx context.Context, id string) ([]Short, error) {
	var shorts []Short
	err := r.db.With(ctx).Select().From("shorts").Where(dbx.HashExp{"user_id": id}).All(&shorts)
	if err != nil {
		return nil, err
	}

	return shorts, err
}
