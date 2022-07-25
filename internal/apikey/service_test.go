package apikey

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/stretchr/testify/assert"
)

var errCRUD = errors.New("error crud")

func Test_service_CRUD(t *testing.T) {
	logger, _ := log.NewForTest()
	s := NewService(&mockRepository{}, logger, 1)

	ctx := context.Background()
	c := auth.WithUser(ctx, "100", "test@test.io")

	// initial count
	count, _ := s.Count(c)
	assert.Equal(t, 0, count)

	// successful creation
	apikey, err := s.Create(c)
	assert.Nil(t, err)
	assert.NotEmpty(t, apikey.Key)
	key := apikey.Key
	assert.NotEmpty(t, apikey.CreatedAt)
	count, _ = s.Count(c)
	assert.Equal(t, 1, count)

	// get
	_, err = s.Get(ctx, "none")
	assert.NotNil(t, err)
	apiKey, err := s.Get(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, key, apiKey.Key)

	// delete
	_, err = s.Delete(ctx, "none")
	assert.NotNil(t, err)
	apiKey, err = s.Delete(ctx, key)
	assert.Nil(t, err)
	assert.Equal(t, key, apiKey.Key)
	count, _ = s.Count(ctx)
	fmt.Printf("Count: %+v\n", count)
	assert.Equal(t, 1, count)
}

type mockRepository struct {
	items []entity.APIKey
}

func (m mockRepository) Get(ctx context.Context, key string) (entity.APIKey, error) {
	for _, item := range m.items {
		if item.Key == key {
			return item, nil
		}
	}
	return entity.APIKey{}, sql.ErrNoRows
}

func (m mockRepository) GetOwned(ctx context.Context, userID string) ([]APIKey, error) {
	var owned []APIKey
	for _, item := range m.items {
		if item.UserID == userID {
			owned = append(owned, APIKey{
				APIKeyResponse: APIKeyResponse{
					Key:       item.Key,
					CreatedAt: item.CreatedAt,
				},
			})
		}
	}
	return []APIKey{}, sql.ErrNoRows
}

func (m mockRepository) Count(ctx context.Context) (int, error) {
	return len(m.items), nil
}

func (m mockRepository) CountByOwner(ctx context.Context, userID string) (int, error) {
	var count int

	for _, item := range m.items {
		if item.UserID == userID {
			count++
		}
	}

	return count, nil
}

func (m *mockRepository) Create(ctx context.Context, apiKey entity.APIKey) error {
	if apiKey.Key == "error" {
		return errCRUD
	}
	m.items = append(m.items, apiKey)
	return nil
}

func (m *mockRepository) Update(ctx context.Context, apiKey entity.APIKey) error {
	if apiKey.Key == "error" {
		return errCRUD
	}
	for i, item := range m.items {
		if item.Key == apiKey.Key {
			m.items[i] = apiKey
			break
		}
	}
	return nil
}

func (m *mockRepository) Delete(ctx context.Context, key string) error {
	for i, item := range m.items {
		if item.Key == key {
			m.items[i] = m.items[len(m.items)-1]
			m.items = m.items[:len(m.items)-1]
			break
		}
	}
	return nil
}
