package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/pkg/dbcontext"
	"github.com/garaekz/goshort/pkg/log"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"golang.org/x/crypto/bcrypt"
)

// Service encapsulates the authentication logic.
type Service interface {
	// Login authenticate authenticates a user using email and password.
	// It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	Login(ctx context.Context, email, password string) (string, error)
}

// Identity represents an authenticated user identity.
type Identity interface {
	// GetID returns the user ID.
	GetID() string
	// GetEmail returns the user name.
	GetEmail() string
}

type service struct {
	db              *dbcontext.DB
	signingKey      string
	tokenExpiration int
	logger          log.Logger
}

// NewService creates a new authentication service.
func NewService(db *dbcontext.DB, signingKey string, tokenExpiration int, logger log.Logger) Service {
	return service{db, signingKey, tokenExpiration, logger}
}

// Login authenticates a user and generates a JWT token if authentication succeeds.
// Otherwise, an error is returned.
func (s service) Login(ctx context.Context, email, password string) (string, error) {
	if identity := s.authenticate(ctx, email, password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unauthorized("")
}

// authenticate authenticates a user using email and password.
// If email and password are correct, an identity is returned. Otherwise, nil is returned.
func (s service) authenticate(ctx context.Context, email, password string) Identity {
	logger := s.logger.With(ctx, "email", email)
	user := entity.User{}

	if err := s.db.With(ctx).Select().From("users as u").Where(dbx.HashExp{"u.email": email, "u.is_active": true}).One(&user); err != nil {
		fmt.Println(err)
		return nil
	}
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Printf("%s %s\n", password, string(pass))
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println(err)
		logger.Infof("authentication failed")
		return nil
	}

	logger.Infof("authentication successful")
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
