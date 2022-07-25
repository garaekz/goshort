package apikey

import (
	"context"

	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/dbcontext"
	"github.com/garaekz/goshort/pkg/log"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access apikeys from the data source.
type Repository interface {
	// Get returns the apikey with the specified apikey Key.
	Get(ctx context.Context, key string) (entity.APIKey, error)
	// GetOwned returns the apikeys of the specified user.
	GetOwned(ctx context.Context, userID string) ([]APIKey, error)
	// Count returns the number of apikeys.
	Count(ctx context.Context) (int, error)
	// CountByOwner returns the number of apikeys by owner.
	CountByOwner(ctx context.Context, userID string) (int, error)
	// Create saves a new apikey in the storage.
	Create(ctx context.Context, apikey entity.APIKey) error
	// Delete removes the apikey with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// StdLen sets the initial length of the generated code
const StdLen = 4

// apiKeyChars contains the valid characters to use in a code
var apiKeyChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// repository persists shorts in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new API Key repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the short with the specified ID from the database.
func (r repository) Get(ctx context.Context, key string) (entity.APIKey, error) {
	var apiKey entity.APIKey
	err := r.db.With(ctx).Select().Model(key, &apiKey)
	if err != nil {
		return apiKey, err
	}
	return apiKey, err
}

// Get reads the short with the specified ID from the database.
func (r repository) GetOwned(ctx context.Context, userID string) ([]APIKey, error) {
	var apiKeys []APIKey
	err := r.db.With(ctx).Select().From("keys").Where(dbx.HashExp{"user_id": userID}).All(&apiKeys)
	if err != nil {
		return apiKeys, err
	}
	return apiKeys, err
}

// Create saves a new short record in the database.
// It returns the ID of the newly inserted short record.
func (r repository) Create(ctx context.Context, apiKey entity.APIKey) error {
	return r.db.With(ctx).Model(&apiKey).Insert()
}

// Delete deletes an short with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	short, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&short).Delete()
}

// Count returns the number of the short records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("keys").Row(&count)
	return count, err
}

// CountByOwner returns the number of the short records in the database.
func (r repository) CountByOwner(ctx context.Context, userID string) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("keys").Where(dbx.HashExp{"user_id": userID}).Row(&count)
	return count, err
}
