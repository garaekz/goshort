package user

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/dbcontext"
	"github.com/garaekz/goshort/pkg/log"
)

// Repository encapsulates the logic to access users from the data source.
type Repository interface {
	// Get returns the user with the specified user ID.
	Get(ctx context.Context, id string) (entity.User, error)
	// Count returns the number of users.
	Count(ctx context.Context) (int, error)
	// Query returns the list of users with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.User, error)
	// Create saves a new user in the storage.
	Create(ctx context.Context, user entity.User) error
	// Update updates the user with given ID in the storage.
	Update(ctx context.Context, user entity.User) error
	// Delete removes the user with given ID from the storage.
	Delete(ctx context.Context, id string) error
	// Link returns the links related to the specified user ID.
	Links(ctx context.Context, offset, limit int) ([]entity.Link, error)
}

// repository persists users in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new user repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the user with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().Model(id, &user)
	return user, err
}

// Create saves a new user record in the database.
// It returns the ID of the newly inserted user record.
func (r repository) Create(ctx context.Context, user entity.User) error {
	return r.db.With(ctx).Model(&user).Exclude("IsActive").Insert()
}

// Update saves the changes to an user in the database.
func (r repository) Update(ctx context.Context, user entity.User) error {
	return r.db.With(ctx).Model(&user).Update()
}

// Delete deletes an user with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	user, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&user).Delete()
}

// Count returns the number of the user records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("users").Row(&count)
	return count, err
}

// Query retrieves the user records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.User, error) {
	var users []entity.User
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&users)
	return users, err
}

// Links reads the user ID and searchs for the links related on the database.
func (r repository) Links(ctx context.Context, offset, limit int) ([]entity.Link, error) {
	var links []entity.Link
	identity := auth.CurrentUser(ctx)

	err := r.db.With(ctx).
		Select().
		Where(dbx.HashExp{"user_id": identity.GetID()}).
		OrderBy("created_at desc").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&links)
	return links, err
}
