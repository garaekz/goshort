package auth

import (
	"context"
	"testing"

	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/internal/test"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "users")
	repo := NewRepository(db, logger)

	ctx := context.Background()

	// register new user
	user := entity.User{
		ID:       "100",
		Email:    "test1@test.io",
		Password: "test",
		IsActive: true,
	}
	err := repo.Register(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, "test1@test.io", user.Email)
	err = repo.Register(ctx, user)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "pq: duplicate key value violates unique constraint \"users_pkey\"")

	// get user by email
	user, err = repo.GetUserByEmail(ctx, "test1@test.io")
	// assert.Nil(t, err)
	// assert.Equal(t, "test1@test.io", user.Email)

	// get me
	// short, err := repo.Get(ctx, "100")
	// assert.Nil(t, err)
	// assert.Equal(t, "test1@test.io", short.Email)
	// _, err = repo.Get(ctx, "test0")
	// assert.Equal(t, sql.ErrNoRows, err)
}
