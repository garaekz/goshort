package apikey

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
	repo := &mockRepository{items: []entity.APIKey{
		{Key: "123", UserID: "100", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsActive: true},
		{Key: "963", UserID: "200", CreatedAt: time.Now(), UpdatedAt: time.Now(), IsActive: true},
	}}
	RegisterHandlers(router.Group(""), NewService(repo, logger, 2), auth.MockAuthHandler, logger)
	header := auth.MockAuthHeader()

	tests := []test.APITestCase{
		{Name: "get 123", Method: "GET", URL: "/apikeys/123", Body: "", Header: header, WantStatus: http.StatusOK, WantResponse: `*123*`},
		{Name: "get unknown", Method: "GET", URL: "/apikeys/1234", Body: "", Header: header, WantStatus: http.StatusNotFound, WantResponse: ""},
		{Name: "create ok", Method: "POST", URL: "/apikeys", Body: "", Header: header, WantStatus: http.StatusCreated, WantResponse: "*key*"},
		{Name: "create auth error", Method: "POST", URL: "/apikeys", Body: `{"url":"test"}`, Header: nil, WantStatus: http.StatusUnauthorized, WantResponse: ""},
		{Name: "delete ok", Method: "DELETE", URL: "/apikeys/123", Body: ``, Header: header, WantStatus: http.StatusOK, WantResponse: "*123*"},
		{Name: "delete verify", Method: "DELETE", URL: "/apikeys/123", Body: ``, Header: header, WantStatus: http.StatusNotFound, WantResponse: ""},
		{Name: "delete auth error", Method: "DELETE", URL: "/apikeys/123", Body: ``, Header: nil, WantStatus: http.StatusUnauthorized, WantResponse: ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
