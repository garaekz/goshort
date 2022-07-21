package short

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
	RegisterHandlers(router.Group(""), NewService(repo, logger), auth.MockAuthHandler, auth.MockAuthHandler, logger)
	header := auth.MockAuthHeader()

	tests := []test.APITestCase{
		{Name: "get all", Method: "GET", URL: "/shorts", Body: "", Header: header, WantStatus: http.StatusOK, WantResponse: `*"total_count":1*`},
		{Name: "get 123", Method: "GET", URL: "/shorts/123", Body: "", Header: header, WantStatus: http.StatusOK, WantResponse: `*short123*`},
		{Name: "get unknown", Method: "GET", URL: "/shorts/1234", Body: "", Header: header, WantStatus: http.StatusNotFound, WantResponse: ""},
		{Name: "create ok", Method: "POST", URL: "/shorts", Body: `{"url":"test"}`, Header: header, WantStatus: http.StatusCreated, WantResponse: "*test*"},
		{Name: "create ok count", Method: "GET", URL: "/shorts", Body: "", Header: header, WantStatus: http.StatusOK, WantResponse: `*"total_count":2*`},
		{Name: "create auth error", Method: "POST", URL: "/shorts", Body: `{"url":"test"}`, Header: nil, WantStatus: http.StatusUnauthorized, WantResponse: ""},
		{Name: "create input error", Method: "POST", URL: "/shorts", Body: `"name":"test"}`, Header: header, WantStatus: http.StatusBadRequest, WantResponse: ""},
		{Name: "update ok", Method: "PUT", URL: "/shorts/123", Body: `{"url":"shortxyz"}`, Header: header, WantStatus: http.StatusOK, WantResponse: "*shortxyz*"},
		{Name: "update verify", Method: "GET", URL: "/shorts/123", Body: "", Header: header, WantStatus: http.StatusOK, WantResponse: `*shortxyz*`},
		{Name: "update auth error", Method: "PUT", URL: "/shorts/123", Body: `{"name":"shortxyz"}`, Header: nil, WantStatus: http.StatusUnauthorized, WantResponse: ""},
		{Name: "update input error", Method: "PUT", URL: "/shorts/123", Body: `"name":"shortxyz"}`, Header: header, WantStatus: http.StatusBadRequest, WantResponse: ""},
		{Name: "delete ok", Method: "DELETE", URL: "/shorts/123", Body: ``, Header: header, WantStatus: http.StatusOK, WantResponse: "*shortxyz*"},
		{Name: "delete verify", Method: "DELETE", URL: "/shorts/123", Body: ``, Header: header, WantStatus: http.StatusNotFound, WantResponse: ""},
		{Name: "delete auth error", Method: "DELETE", URL: "/shorts/123", Body: ``, Header: nil, WantStatus: http.StatusUnauthorized, WantResponse: ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
