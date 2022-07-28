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

func (m mockRepository) Get(_ context.Context, id string) (entity.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return user, nil
		}
	}

	return entity.User{}, sql.ErrNoRows
}

func (apiKeyMockRepository) Get(_ context.Context, _ string) (entity.APIKey, error) {
	return entity.APIKey{}, sql.ErrNoRows
}

func (apiKeyMockRepository) GetOwned(_ context.Context, _ string) ([]apikey.APIKey, error) {
	return []apikey.APIKey{}, sql.ErrNoRows
}

func (apiKeyMockRepository) Count(_ context.Context) (int, error) {
	return 0, nil
}

func (apiKeyMockRepository) CountByOwner(_ context.Context, _ string) (int, error) {
	var count int
	return count, nil
}

func (apiKeyMockRepository) Create(_ context.Context, _ entity.APIKey) error {
	return nil
}

func (apiKeyMockRepository) Update(_ context.Context, _ entity.APIKey) error {
	return nil
}

func (apiKeyMockRepository) Delete(_ context.Context, _ string) error {
	return nil
}

func (shortMockRepository) Get(_ context.Context, _ string) (entity.Short, error) {
	return entity.Short{}, sql.ErrNoRows
}

func (shortMockRepository) Count(_ context.Context) (int, error) {
	return 0, nil
}

func (m shortMockRepository) Query(_ context.Context, _, _ int) ([]entity.Short, error) {
	return m.shorts, nil
}

func (shortMockRepository) Create(_ context.Context, _ entity.Short) error {
	return nil
}

func (shortMockRepository) Update(_ context.Context, _ entity.Short) error {
	return nil
}

func (shortMockRepository) Delete(_ context.Context, _ string) error {
	return nil
}

func (shortMockRepository) GenerateUniqueCode(_ context.Context) (string, error) {
	return "code", nil
}

func (shortMockRepository) GetByOriginalURL(_ context.Context, _, _ string) (entity.Short, error) {
	return entity.Short{}, nil
}

func (m *shortMockRepository) GetOwned(_ context.Context, _ string) ([]entity.Short, error) {
	return m.shorts, nil
}
