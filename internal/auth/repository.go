package auth

import (
	"context"

	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/dbcontext"
	"github.com/garaekz/goshort/pkg/log"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Repository encapsulates the logic to access info from the data source.
type Repository interface {
	// GetUserByEmail passes the email to the database and returns the user
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	// GetUserByAPIKey passes the api key to the database and returns the user
	GetUserByAPIKey(ctx context.Context, apiKey string) (entity.User, error)
	// Register saves a new user in the database.
	Register(ctx context.Context, user entity.User) error
	// CreteEmailVerification creates new a record in the database for email validation
	CreateEmailVerification(ctx context.Context, validation entity.EmailVerification) error
	// VerifyEmail verifies a user's email, updates user and deletes verification record.
	VerifyEmail(ctx context.Context, validation VerifyRequest) error
}

// repository persists users in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new auth repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// GetUserByEmail passes the email to the database and returns the user
func (r repository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).Select().From("users").Where(dbx.HashExp{"email": email, "is_active": true}).One(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUserByAPIKey passes the api key to the database and returns the user
func (r repository) GetUserByAPIKey(ctx context.Context, apiKey string) (entity.User, error) {
	var user entity.User
	err := r.db.With(ctx).
		Select("users.*").
		From("users").
		LeftJoin("keys", dbx.NewExp("users.id=keys.user_id")).
		Where(dbx.HashExp{"keys.key": apiKey}).
		One(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

// Register saves a new user in the database.
func (r repository) Register(ctx context.Context, user entity.User) error {
	return r.db.With(ctx).Model(&user).Insert()
}

func (r repository) CreateEmailVerification(ctx context.Context, validation entity.EmailVerification) error {
	return r.db.With(ctx).Model(&validation).Insert()
}

func (r repository) VerifyEmail(ctx context.Context, validation VerifyRequest) error {
	var user entity.User

	err := r.db.DB().Transactional(func(tx *dbx.Tx) error {
		if err := tx.Select().From("users").Where(dbx.HashExp{"user_id": validation.UserID}).One(&user); err != nil {
			return err
		}

		if _, err := tx.Update("users", dbx.Params{"email_verified": true}, dbx.HashExp{"user_id": validation.UserID}).Execute(); err != nil {
			return err
		}

		if _, err := tx.Delete("email_verifications", dbx.HashExp{"user_id": validation.UserID}).Execute(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
