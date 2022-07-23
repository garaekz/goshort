package my

import (
	"net/http"
	"testing"
	"time"

	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/entity"
	"github.com/garaekz/goshort/internal/test"
	"github.com/garaekz/goshort/pkg/log"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	repo := &mockRepository{items: []entity.Short{
		{Code: "123", OriginalURL: "short123.com", Visits: 0, UserID: "1", CreatorIP: "127.0.0.1", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: nil},
	}}
	RegisterHandlers(router.Group(""), NewService(repo, logger), auth.MockAuthHandler, logger)
	header := auth.MockAuthHeader()

	tests := []test.APITestCase{
		{Name: "get my user", Method: "GET", URL: "/me", Body: "", Header: header, WantStatus: http.StatusOK, WantResponse: `*"email":"`},
		{Name: "get my user auth error", Method: "GET", URL: "/me", Body: "", Header: nil, WantStatus: http.StatusUnauthorized, WantResponse: ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
