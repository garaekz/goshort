package link

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/dbcontext"
	"github.com/garaekz/goshort/pkg/log"
)

// Repository encapsulates the logic to access links from the data source.
type Repository interface {
	// Get returns the link with the specified link ID.
	Get(ctx context.Context, id string) (entity.Link, error)
	// Count returns the number of links.
	Count(ctx context.Context) (int, error)
	// Query returns the list of links with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Link, error)
	// Create saves a new link in the storage.
	Create(ctx context.Context, link entity.Link) error
	// Update updates the link with given ID in the storage.
	Update(ctx context.Context, link entity.Link) error
	// Delete removes the link with given ID from the storage.
	Delete(ctx context.Context, id string) error
	// GetByOriginalURL checks for the existence of a record by a given URL string.
	GetByOriginalURL(ctx context.Context, URL string) (*entity.Link, error)
	// GetByCode checks for the existence of a record by a given code string.
	GetByCode(ctx context.Context, code string) (*entity.Link, error)
	// GenerateUniqueCode checks for the existence of a record by a random code and returns the code if nothing found.
	GenerateUniqueCode(ctx context.Context) (string, error)
}

// StdLen sets the initial length of the generated code
const (
	StdLen = 4
)

// StdChars contains the valid characters to use in a code
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// repository persists links in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new link repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the link with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Link, error) {
	var link entity.Link
	err := r.db.With(ctx).Select().Model(id, &link)
	return link, err
}

// Create saves a new link record in the database.
// It returns the ID of the newly inserted link record.
func (r repository) Create(ctx context.Context, link entity.Link) error {
	return r.db.With(ctx).Model(&link).Insert()
}

// Update saves the changes to an link in the database.
func (r repository) Update(ctx context.Context, link entity.Link) error {
	return r.db.With(ctx).Model(&link).Update()
}

// Delete deletes an link with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	link, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&link).Delete()
}

// Count returns the number of the link records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("urls").Row(&count)
	return count, err
}

// Query retrieves the link records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Link, error) {
	var links []entity.Link
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&links)
	return links, err
}

// GetByOriginalURL retrieves a single link given a OriginalURL string.
func (r repository) GetByOriginalURL(ctx context.Context, URL string) (*entity.Link, error) {
	var link *entity.Link

	err := r.db.With(ctx).
		Select().
		From("urls").
		Where(dbx.HashExp{"original_url": URL}).
		One(&link)

	return link, err
}

// GetByCode retrieves a single link given a Code string.
func (r repository) GetByCode(ctx context.Context, code string) (*entity.Link, error) {
	var link *entity.Link

	err := r.db.With(ctx).
		Select().
		From("urls").
		Where(dbx.HashExp{"code": code}).
		One(&link)

	return link, err
}

// GenerateUniqueCode checks for the existence of a record by a random code and returns the code if nothing found.
func (r repository) GenerateUniqueCode(ctx context.Context) (string, error) {
	n := StdLen
	i := 0

	for {
		code, err := randomCode(n, StdChars)

		if err != nil {
			return "", err
		}

		_, err = r.GetByCode(ctx, code)

		if err != nil {
			if err.Error() == "sql: no rows in result set" {
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
