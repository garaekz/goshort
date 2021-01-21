package page

import (
	"context"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/dbcontext"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Repository encapsulates the logic to access pages from the data source.
type Repository interface {
	// Get returns the page with the specified page ID.
	GetByCode(ctx context.Context, code string) (*entity.Link, error)
}

// repository persists pages in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new page repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
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
