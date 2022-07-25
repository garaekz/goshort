package my

import (
	"context"
	"database/sql"
	"testing"

	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/internal/test"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "shorts")
	test.ResetTables(t, db, "users")
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

	// get me
	short, err := repo.Get(ctx, "100")
	assert.Nil(t, err)
	assert.Equal(t, "test1@test.io", short.Email)
	_, err = repo.Get(ctx, "test0")
	assert.Equal(t, sql.ErrNoRows, err)
}
