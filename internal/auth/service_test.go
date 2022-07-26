package auth

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	items []entity.User
	keys  []struct {
		Key    string
		UserID string
	}
}

func Test_service_Authenticate(t *testing.T) {
	logger, _ := log.NewForTest()
	repo := &mockRepository{
		items: []entity.User{
			{ID: "100", Email: "test@test.io", Password: "pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsActive: true},
		},
		keys: []struct {
			Key    string
			UserID string
		}{
			{Key: "9876543210", UserID: "rrr"},
			{Key: "0123456789", UserID: "100"},
		},
	}
	s := NewService(repo, "test", 100, logger)
	_, err := s.Login(context.Background(), "unknown", "bad")
	assert.Equal(t, errors.Unauthorized("Login failed, please check your credentials"), err)
	token, err := s.Login(context.Background(), "test@test.io", "pass")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func Test_service_authenticate(t *testing.T) {
	logger, _ := log.NewForTest()
	repo := &mockRepository{
		items: []entity.User{
			{ID: "100", Email: "test@test.io", Password: "pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsActive: true},
		},
		keys: []struct {
			Key    string
			UserID string
		}{
			{Key: "9876543210", UserID: "rrr"},
			{Key: "0123456789", UserID: "100"},
		},
	}
	s := service{repo, "test", 100, logger}
	assert.Nil(t, s.authenticate(context.Background(), "unknown", "bad"))
	assert.NotNil(t, s.authenticate(context.Background(), "test@test.io", "pass"))
}

func Test_service_GenerateJWT(t *testing.T) {
	logger, _ := log.NewForTest()
	repo := &mockRepository{
		items: []entity.User{
			{ID: "100", Email: "test@test.io", Password: "pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsActive: true},
		},
		keys: []struct {
			Key    string
			UserID string
		}{
			{Key: "9876543210", UserID: "rrr"},
			{Key: "0123456789", UserID: "100"},
		},
	}
	s := service{repo, "test", 100, logger}
	token, err := s.generateJWT(entity.User{
		ID:    "100",
		Email: "test@test.io",
	})
	if assert.Nil(t, err) {
		assert.NotEmpty(t, token)
	}
}

func (m mockRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	for _, item := range m.items {
		if item.Email == email {
			return item, nil
		}
	}
	return entity.User{}, sql.ErrNoRows
}

func (m mockRepository) GetUserByAPIKey(ctx context.Context, apiKey string) (entity.User, error) {
	var userID string
	for _, key := range m.keys {
		if key.UserID == apiKey {
			userID = key.UserID
		}
	}
	for _, item := range m.items {
		if item.ID == userID {
			return item, nil
		}
	}
	return entity.User{}, sql.ErrNoRows
}

func (m mockRepository) Register(ctx context.Context, user entity.User) error {
	if user.ID == "error" {
		return errors.Unauthorized("")
	}
	m.items = append(m.items, user)
	return nil
}
