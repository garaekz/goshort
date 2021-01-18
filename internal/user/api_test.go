package user

import (
	"github.com/qiangxue/go-rest-api/internal/auth"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/test"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"net/http"
	"testing"
	"time"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	repo := &mockRepository{items: []entity.User{
		{"123", "user@mail.io", "dummypass", time.Now(), time.Now()},
	}}
	RegisterHandlers(router.Group(""), NewService(repo, logger), auth.MockAuthHandler, logger)
	header := auth.MockAuthHeader()

	tests := []test.APITestCase{
		{"get all", "GET", "/users", "", nil, http.StatusOK, `*"total_count":1*`},
		{"get 123", "GET", "/users/123", "", nil, http.StatusOK, `*user123*`},
		{"get unknown", "GET", "/users/1234", "", nil, http.StatusNotFound, ""},
		{"create ok", "POST", "/users", `{"email":"newmail@mail.io", "password":"dummypass"}`, header, http.StatusCreated, "*test*"},
		{"create ok count", "GET", "/users", "", nil, http.StatusOK, `*"total_count":2*`},
		{"create auth error", "POST", "/users", `{"name":"test"}`, nil, http.StatusUnauthorized, ""},
		{"create input error", "POST", "/users", `"name":"test"}`, header, http.StatusBadRequest, ""},
		{"update ok", "PUT", "/users/123", `{"name":"albumxyz"}`, header, http.StatusOK, "*albumxyz*"},
		{"update verify", "GET", "/users/123", "", nil, http.StatusOK, `*albumxyz*`},
		{"update auth error", "PUT", "/users/123", `{"name":"albumxyz"}`, nil, http.StatusUnauthorized, ""},
		{"update input error", "PUT", "/users/123", `"name":"albumxyz"}`, header, http.StatusBadRequest, ""},
		{"delete ok", "DELETE", "/users/123", ``, header, http.StatusOK, "*albumxyz*"},
		{"delete verify", "DELETE", "/users/123", ``, header, http.StatusNotFound, ""},
		{"delete auth error", "DELETE", "/users/123", ``, nil, http.StatusUnauthorized, ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
