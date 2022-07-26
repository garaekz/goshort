package short

import (
	"context"
	"database/sql"

	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/dbcontext"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/garaekz/goshort/pkg/utils"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access shorts from the data source.
type Repository interface {
	// Get returns the short with the specified short ID.
	Get(ctx context.Context, id string) (entity.Short, error)
	// GetOwned returns the short with the specified user ID.
	GetOwned(ctx context.Context, userID string) ([]entity.Short, error)
	// Count returns the number of shorts.
	Count(ctx context.Context) (int, error)
	// Query returns the list of shorts with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Short, error)
	// Create saves a new short in the storage.
	Create(ctx context.Context, short entity.Short) error
	// Update updates the short with given ID in the storage.
	Update(ctx context.Context, short entity.Short) error
	// Delete removes the short with given ID from the storage.
	Delete(ctx context.Context, id string) error
	// GetByOriginalURL checks for the existence of a record by a given URL string.
	GetByOriginalURL(ctx context.Context, URL, userID string) (entity.Short, error)
	// GenerateUniqueCode checks for the existence of a record by a random code and returns the code if nothing found.
	GenerateUniqueCode(ctx context.Context) (string, error)
}

// StdLen sets the initial length of the generated code
const StdLen = 4

// StdChars contains the valid characters to use in a code
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-*+#@")

// repository persists shorts in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new short repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the short with the specified ID from the database.
func (r repository) Get(ctx context.Context, code string) (entity.Short, error) {
	var short entity.Short
	err := r.db.With(ctx).Select().Model(code, &short)
	return short, err
}

// GetOwned reads the short with the specified user ID from the database.
func (r repository) GetOwned(ctx context.Context, userID string) ([]entity.Short, error) {
	var shorts []entity.Short
	err := r.db.With(ctx).Select().From("shorts").Where(dbx.HashExp{"user_id": userID}).All(&shorts)
	return shorts, err
}

// Create saves a new short record in the database.
// It returns the ID of the newly inserted short record.
func (r repository) Create(ctx context.Context, short entity.Short) error {
	return r.db.With(ctx).Model(&short).Insert()
}

// Update saves the changes to an short in the database.
func (r repository) Update(ctx context.Context, short entity.Short) error {
	return r.db.With(ctx).Model(&short).Update()
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
	err := r.db.With(ctx).Select("COUNT(*)").From("shorts").Row(&count)
	return count, err
}

// Query retrieves the short records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Short, error) {
	var shorts []entity.Short
	err := r.db.With(ctx).
		Select().
		OrderBy("created_at desc").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&shorts)
	return shorts, err
}

// GetByOriginalURL retrieves a single link given a OriginalURL string.
func (r repository) GetByOriginalURL(ctx context.Context, URL, userID string) (entity.Short, error) {
	var short entity.Short

	err := r.db.With(ctx).
		Select().
		From("shorts").
		Where(dbx.HashExp{"original_url": URL, "user_id": userID}).
		One(&short)

	return short, err
}

// GenerateUniqueCode checks for the existence of a record by a random code and returns the code if nothing found.
func (r repository) GenerateUniqueCode(ctx context.Context) (string, error) {
	n := StdLen
	i := 0

	for {
		code, err := utils.RandomString(n, StdChars)
		if err != nil {
			return "", err
		}

		_, err = r.Get(ctx, code)
		if err != nil {
			if err == sql.ErrNoRows {
				return code, nil
			}
			return "", err
		}
		if i%10 == 0 {
			n++
		}
		i++
	}
}
