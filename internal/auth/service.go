package auth

import (
	"context"
	defaultErrors "errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

// Service encapsulates the authentication logic.
type Service interface {
	// authenticate authenticates a user using username and password.
	// It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	Login(ctx context.Context, email, password string) (string, error)
	// Register creates a new user and authenticates it.
	Register(ctx context.Context, email, password string) (string, error)
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

// NewService creates a new authentication service.
func NewService(repo Repository, signingKey string, tokenExpiration int, logger log.Logger) Service {
	return service{repo, signingKey, tokenExpiration, logger}
}

var errRegister = defaultErrors.New(`pq: duplicate key value violates unique constraint "users_email_key"`)

// Login authenticates a user and generates a JWT token if authentication succeeds.
// Otherwise, an error is returned.
func (s service) Login(ctx context.Context, email, password string) (string, error) {
	if identity := s.authenticate(ctx, email, password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unauthorized("Login failed, please check your credentials")
}

// Register creates a new user and authenticates it.
func (s service) Register(ctx context.Context, email, password string) (string, error) {
	userData := entity.User{
		ID:        entity.GenerateID(),
		Email:     email,
		Password:  password,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.repo.Register(ctx, userData)
	if err != nil {
		if err.Error() == errRegister.Error() {
			return "", errors.UserAlreadyExists("The user you're trying to register already exists")
		}
		return "", err
	}

	return s.generateJWT(userData)
}

/*
Method authenticate validate a user using username and password.
If username and password are correct, an identity is returned. Otherwise, nil is returned.
*/
func (s service) authenticate(ctx context.Context, email, password string) Identity {
	logger := s.logger.With(ctx, "email", email)

	// This shall only work on test environment.
	if s.signingKey == "test" && email == "test@test.io" && password == "pass" {
		logger.Infof("authentication successful")
		return entity.User{ID: "100", Email: "test@test.io"}
	}

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		logger.Infof("User not found: Authentication failed")
		return nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		logger.Infof("Authentication failed")
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
