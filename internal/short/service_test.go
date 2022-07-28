package short

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/stretchr/testify/assert"
)

var errCRUD = errors.New("error crud")
var errParse = errors.New("URL is not in a valid format")

func TestCreateShortRequest_Validate(t *testing.T) {
	tests := []struct {
		URL       string
		model     CreateShortRequest
		wantError bool
	}{
		{"required", CreateShortRequest{URL: ""}, true},
		{"too long", CreateShortRequest{URL: "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.URL, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func TestUpdateShortRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		model     UpdateShortRequest
		wantError bool
	}{
		{"success", UpdateShortRequest{URL: "test"}, false},
		{"required", UpdateShortRequest{URL: ""}, true},
		{"too long", UpdateShortRequest{URL: "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func Test_service_CRUD(t *testing.T) {
	logger, _ := log.NewForTest()
	s := NewService(&mockRepository{}, logger)

	ctx := context.Background()
	c := auth.WithUser(ctx, "100", "test@test.io")
	// initial count
	count, _ := s.Count(c)
	assert.Equal(t, 0, count)

	// successful creation
	short, err := s.Create(c, CreateShortRequest{URL: "test.io", IP: "127.0.0.1"})
	assert.Nil(t, err)
	assert.NotEmpty(t, short.Code)
	code := short.Code
	assert.Equal(t, "http://test.io", short.OriginalURL)
	assert.NotEmpty(t, short.CreatedAt)
	assert.NotEmpty(t, short.UpdatedAt)
	count, _ = s.Count(c)
	assert.Equal(t, 1, count)

	// validation error in creation
	_, err = s.Create(c, CreateShortRequest{URL: ""})
	assert.NotNil(t, err)
	count, _ = s.Count(c)
	assert.Equal(t, 1, count)

	// parse url error in creation
	_, err = s.Create(c, CreateShortRequest{URL: "```", IP: "127.0.0.1"})
	assert.Equal(t, errParse, err)
	count, _ = s.Count(c)
	assert.Equal(t, 1, count)

	_, _ = s.Create(c, CreateShortRequest{URL: "test2", IP: "127.0.0.1"})

	// update
	short, err = s.Update(ctx, code, UpdateShortRequest{URL: "https://updated.io"})
	assert.Nil(t, err)
	assert.Equal(t, "https://updated.io", short.OriginalURL)
	_, err = s.Update(ctx, "none", UpdateShortRequest{URL: "test updated"})
	assert.NotNil(t, err)

	// validation error in update
	_, err = s.Update(ctx, code, UpdateShortRequest{URL: ""})
	assert.NotNil(t, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 2, count)

	// unexpected error in update
	_, err = s.Update(ctx, code, UpdateShortRequest{URL: "```"})
	assert.Equal(t, errParse, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 2, count)

	// get
	_, err = s.Get(ctx, "none")
	assert.NotNil(t, err)
	short, err = s.Get(ctx, code)
	assert.Nil(t, err)
	assert.Equal(t, "https://updated.io", short.OriginalURL)
	assert.Equal(t, code, short.Code)

	// query
	shorts, _ := s.Query(ctx, 0, 0)
	assert.Equal(t, 2, len(shorts))

	// delete
	_, err = s.Delete(ctx, "none")
	assert.NotNil(t, err)
	short, err = s.Delete(ctx, code)
	assert.Nil(t, err)
	assert.Equal(t, code, short.Code)
	count, _ = s.Count(ctx)
	assert.Equal(t, 1, count)
}

type mockRepository struct {
	items []entity.Short
}

func (m mockRepository) Get(_ context.Context, code string) (entity.Short, error) {
	for _, item := range m.items {
		if item.Code == code {
			return item, nil
		}
	}
	return entity.Short{}, sql.ErrNoRows
}

func (m mockRepository) GetOwned(_ context.Context, userID string) ([]entity.Short, error) {
	var shorts []entity.Short
	for _, item := range m.items {
		if item.UserID == userID {
			shorts = append(shorts, item)
		}
	}
	if len(shorts) == 0 {
		return nil, sql.ErrNoRows
	}
	return shorts, nil
}

func (m mockRepository) Count(_ context.Context) (int, error) {
	return len(m.items), nil
}

func (m mockRepository) Query(_ context.Context, _, _ int) ([]entity.Short, error) {
	return m.items, nil
}

func (m *mockRepository) Create(_ context.Context, short entity.Short) error {
	if short.Code == "error" {
		return errCRUD
	}
	m.items = append(m.items, short)
	return nil
}

func (m *mockRepository) Update(_ context.Context, short entity.Short) error {
	if short.Code == "error" {
		return errCRUD
	}
	for i, item := range m.items {
		if item.Code == short.Code {
			m.items[i] = short
			break
		}
	}
	return nil
}

func (m *mockRepository) Delete(_ context.Context, code string) error {
	for i, item := range m.items {
		if item.Code == code {
			m.items[i] = m.items[len(m.items)-1]
			m.items = m.items[:len(m.items)-1]
			break
		}
	}
	return nil
}

func (*mockRepository) GenerateUniqueCode(_ context.Context) (string, error) {
	return "code", nil
}

func (*mockRepository) GetByOriginalURL(_ context.Context, _, _ string) (entity.Short, error) {
	return entity.Short{}, nil
}
