package auth

import (
	"context"
	defaultErrors "errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/garaekz/goshort/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

// Service encapsulates the authentication logic.
type Service interface {
	// authenticate authenticates a user using username and password.
	// It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	Login(ctx context.Context, email, password string) (string, error)
	// Register creates a new user and authenticates it.
	Register(ctx context.Context, email, password string) error
	// Verify verifies a user's email address.
	Verify(ctx context.Context, id, token string) error
}

// Identity represents an authenticated user identity.
type Identity interface {
	// GetID returns the user ID.
	GetID() string
	// GetEmail returns the user name.
	GetEmail() string
}

type service struct {
	repo            Repository
	signingKey      string
	tokenExpiration int
	logger          log.Logger
}

// VerifyRequest represents a request to verify a user's email address.
type VerifyRequest struct {
	UserID string
	Token  string
}

// NewService creates a new authentication service.
func NewService(repo Repository, signingKey string, tokenExpiration int, logger log.Logger) Service {
	return service{repo, signingKey, tokenExpiration, logger}
}

var errRegister = defaultErrors.New(`pq: duplicate key value violates unique constraint "users_email_key"`)

// Login authenticates a user and generates a JWT token if authentication succeeds.
func (s service) Login(ctx context.Context, email, password string) (string, error) {
	if identity := s.authenticate(ctx, email, password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unauthorized("Login failed, please check your credentials")
}

// Register creates a new user and authenticates it.
func (s service) Register(ctx context.Context, email, password string) error {
	var tokenChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-=")
	token, err := utils.RandomString(64, tokenChars)
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	userData := entity.User{
		ID:            entity.GenerateID(),
		Email:         email,
		Password:      string(pass),
		IsActive:      true,
		EmailVerified: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err = s.repo.Register(ctx, userData)
	if err != nil {
		if err.Error() == errRegister.Error() {
			return errors.BadRequest("The user you're trying to register already exists")
		}
		return err
	}

	err = s.repo.CreateEmailVerification(ctx, entity.EmailVerification{
		UserID:    userData.ID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour),
	})
	if err != nil {
		return err
	}

	return nil
}

func (s service) Verify(ctx context.Context, id, token string) error {
	verification, err := s.repo.GetEmailVerification(ctx, id, token)
	if time.Now().After(verification.ExpiresAt) {
		return errors.BadRequest("Your verification code has expired")
	}
	verifyRequest := VerifyRequest{
		UserID: id,
		Token:  token,
	}
	err = s.repo.VerifyEmail(ctx, verifyRequest)
	if err != nil {
		return err
	}
	return nil
}

/*
Method authenticate validate a user using username and password.
If username and password are correct, an identity is returned. Otherwise, nil is returned.
*/
func (s service) authenticate(ctx context.Context, email, password string) Identity {
	logger := s.logger.With(ctx, "email", email)

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		logger.Infof("User not found: Authentication failed")
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logger.Infof("Authentication failed")
		return nil
	}

	if !user.EmailVerified {
		logger.Infof("User not verified: Authentication failed")
		return nil
	}

	logger.Infof("Authentication successful")
	return entity.User{ID: user.GetID(), Email: user.GetEmail()}
}

// generateJWT generates a JWT that encodes an identity.
func (s service) generateJWT(identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    identity.GetID(),
		"email": identity.GetEmail(),
		"exp":   time.Now().Add(time.Duration(s.tokenExpiration) * time.Hour).Unix(),
	}).SignedString([]byte(s.signingKey))
}
