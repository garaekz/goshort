package apikey

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/internal/test"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "keys")
	repo := NewRepository(db, logger)
	authRepo := auth.NewRepository(db, logger)

	ctx := context.Background()

	// create user
	user := entity.User{
		ID:       "100",
		Email:    "test1@test.io",
		Password: "test",
	}
	err := authRepo.Register(ctx, user)

	// initial count
	count, err := repo.Count(ctx)
	assert.Nil(t, err)

	// initial count by user
	countByOwner, err := repo.CountByOwner(ctx, "100")
	assert.Nil(t, err)

	// create
	err = repo.Create(ctx, entity.APIKey{
		Key:       "ReallyLongAndSecretKey",
		UserID:    "100",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	count2, _ := repo.Count(ctx)
	assert.Equal(t, 1, count2-count)

	countByOwner2, _ := repo.CountByOwner(ctx, "100")
	assert.Equal(t, 1, countByOwner2-countByOwner)

	// get
	apiKey, err := repo.Get(ctx, "ReallyLongAndSecretKey")
	assert.Nil(t, err)
	assert.Equal(t, "ReallyLongAndSecretKey", apiKey.Key)
	_, err = repo.Get(ctx, "test0")
	assert.Equal(t, sql.ErrNoRows, err)

	// delete
	err = repo.Delete(ctx, "ReallyLongAndSecretKey")
	assert.Nil(t, err)
	_, err = repo.Get(ctx, "ReallyLongAndSecretKey")
	assert.Equal(t, sql.ErrNoRows, err)
	err = repo.Delete(ctx, "ReallyLongAndSecretKey")
	assert.Equal(t, sql.ErrNoRows, err)
}
