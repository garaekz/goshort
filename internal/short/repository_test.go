package short

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/internal/test"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "shorts")
	repo := NewRepository(db, logger)

	ctx := context.Background()

	// initial count
	count, err := repo.Count(ctx)
	assert.Nil(t, err)

	// create
	err = repo.Create(ctx, entity.Short{
		Code:        "test1",
		OriginalURL: "http://test.com",
		Visits:      0,
		UserID:      "100",
		CreatorIP:   "127.0.0.1",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	assert.Nil(t, err)
	count2, _ := repo.Count(ctx)
	assert.Equal(t, 1, count2-count)

	// get
	short, err := repo.Get(ctx, "test1")
	assert.Nil(t, err)
	assert.Equal(t, "http://test.com", short.OriginalURL)
	_, err = repo.Get(ctx, "test0")
	assert.Equal(t, sql.ErrNoRows, err)

	// update
	err = repo.Update(ctx, entity.Short{
		Code:        "test1 updated",
		OriginalURL: "http://test.com updated",
		Visits:      0,
		UpdatedAt:   time.Now(),
	})
	assert.Nil(t, err)
	short, _ = repo.Get(ctx, "test1")
	assert.Equal(t, "http://test.com", short.OriginalURL)

	// query
	shorts, err := repo.Query(ctx, 0, count2)
	assert.Nil(t, err)
	assert.Equal(t, count2, len(shorts))

	// delete
	err = repo.Delete(ctx, "test1")
	assert.Nil(t, err)
	_, err = repo.Get(ctx, "test1")
	assert.Equal(t, sql.ErrNoRows, err)
	err = repo.Delete(ctx, "test1")
	assert.Equal(t, sql.ErrNoRows, err)
}
