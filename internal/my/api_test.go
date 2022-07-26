package my

import (
	"net/http"
	"testing"
	"time"

	"github.com/garaekz/goshort/internal/apikey"
	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/internal/short"
	"github.com/garaekz/goshort/internal/test"
	"github.com/garaekz/goshort/pkg/log"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	repo := &mockRepository{
		users: []entity.User{
			{ID: "100", Email: "test@test.io", Password: "pass", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsActive: true},
		},
	}

	keyRepo := &apiKeyMockRepository{}
	shortRepo := &shortMockRepository{}

	RegisterHandlers(router.Group(""), NewService(repo, logger), apikey.NewService(keyRepo, logger, 2), short.NewService(shortRepo, logger), auth.MockAuthHandler, logger)
	header := auth.MockAuthHeader()

	tests := []test.APITestCase{
		{Name: "get my user", Method: "GET", URL: "/me", Body: "", Header: header, WantStatus: http.StatusOK, WantResponse: `*"email":"`},
		{Name: "get my user auth error", Method: "GET", URL: "/me", Body: "", Header: nil, WantStatus: http.StatusUnauthorized, WantResponse: ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
