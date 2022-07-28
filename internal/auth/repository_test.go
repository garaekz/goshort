package auth

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/internal/test"
	"github.com/garaekz/goshort/pkg/dbcontext"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/stretchr/testify/assert"
)

func createsAPIKey(ctx context.Context, db *dbcontext.DB) error {
	apiKey := entity.APIKey{
		UserID: "100",
		Key:    "9876543210",
	}
	return db.With(ctx).Model(&apiKey).Insert()
}

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
	assert.Nil(t, err)
	assert.Equal(t, "test1@test.io", user.Email)
	assert.Equal(t, "100", user.ID)
	user, err = repo.GetUserByEmail(ctx, "none@test.io")
	assert.NotNil(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())

	// get user by api key
	_ = createsAPIKey(ctx, db)
	user, err = repo.GetUserByAPIKey(ctx, "9876543210")
	assert.Nil(t, err)
	assert.Equal(t, "100", user.ID)
	user, err = repo.GetUserByAPIKey(ctx, "0123456789")
	assert.NotNil(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())

	// create email verification
	verification := entity.EmailVerification{
		UserID:    "100",
		Token:     "test",
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}
	verification2 := entity.EmailVerification{
		UserID:    "200",
		Token:     "test2",
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}
	err = repo.CreateEmailVerification(ctx, verification)
	assert.Nil(t, err)
	err = repo.CreateEmailVerification(ctx, verification2)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "pq: insert or update on table \"email_verifications\" violates foreign key constraint \"email_verifications_user_id_fkey\"")

	// verify email
	err = repo.VerifyEmail(ctx, VerifyRequest{
		UserID: "100",
		Token:  "test",
	})

	assert.Nil(t, err)

}
