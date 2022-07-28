package my

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/garaekz/goshort/internal/apikey"
	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/stretchr/testify/assert"
)

var errCRUD = errors.New("error crud")

func Test_service_CRUD(t *testing.T) {
	logger, _ := log.NewForTest()
	s := NewService(&mockRepository{}, logger)

	ctx := context.Background()
	c := auth.WithUser(ctx, "100", "test@test.io")

	_, err := s.GetMyUser(c)
	assert.NotNil(t, err)
}

type mockRepository struct {
	users []entity.User
}

type apiKeyMockRepository struct {
	apiKeys []entity.APIKey
}

type shortMockRepository struct {
	shorts []entity.Short
}

func (m mockRepository) Get(ctx context.Context, id string) (entity.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return user, nil
		}
	}

	return entity.User{}, sql.ErrNoRows
}

func (m apiKeyMockRepository) Get(ctx context.Context, key string) (entity.APIKey, error) {
	return entity.APIKey{}, sql.ErrNoRows
}

func (m apiKeyMockRepository) GetOwned(ctx context.Context, userID string) ([]apikey.APIKey, error) {
	return []apikey.APIKey{}, sql.ErrNoRows
}

func (m apiKeyMockRepository) Count(ctx context.Context) (int, error) {
	return 0, nil
}

func (m apiKeyMockRepository) CountByOwner(ctx context.Context, userID string) (int, error) {
	var count int
	return count, nil
}

func (m *apiKeyMockRepository) Create(ctx context.Context, apiKey entity.APIKey) error {
	return nil
}

func (m *apiKeyMockRepository) Update(ctx context.Context, apiKey entity.APIKey) error {
	return nil
}

func (m *apiKeyMockRepository) Delete(ctx context.Context, key string) error {
	return nil
}

func (m shortMockRepository) Get(ctx context.Context, code string) (entity.Short, error) {
	return entity.Short{}, sql.ErrNoRows
}

func (m shortMockRepository) Count(ctx context.Context) (int, error) {
	return 0, nil
}

func (m shortMockRepository) Query(ctx context.Context, offset, limit int) ([]entity.Short, error) {
	return m.shorts, nil
}

func (m *shortMockRepository) Create(ctx context.Context, short entity.Short) error {
	return nil
}

func (m *shortMockRepository) Update(ctx context.Context, short entity.Short) error {
	return nil
}

func (m *shortMockRepository) Delete(ctx context.Context, code string) error {
	return nil
}

func (m *shortMockRepository) GenerateUniqueCode(ctx context.Context) (string, error) {
	return "code", nil
}

func (m *shortMockRepository) GetByOriginalURL(ctx context.Context, URL, userID string) (entity.Short, error) {
	return entity.Short{}, nil
}

func (m *shortMockRepository) GetOwned(ctx context.Context, userID string) ([]entity.Short, error) {
	return m.shorts, nil
}
