package auth

import (
	"context"
	"net/http"
	"testing"

	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/internal/test"
	"github.com/garaekz/goshort/pkg/log"
)

type mockService struct{}

func (mockService) Login(_ context.Context, email, password string) (string, error) {
	if email == "test@test.io" && password == "pass" {
		return "token-100", nil
	}
	return "", errors.Unauthorized("")
}

func (mockService) Register(_ context.Context, email, password string) error {
	if email == "test@test.io" && password == "pass" {
		return nil
	}
	return errors.Unauthorized("")
}

func (mockService) Verify(_ context.Context, _, _ string) error {
	return nil
}

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)

	RegisterHandlers(router.Group(""), mockService{}, logger)

	tests := []test.APITestCase{
		{Name: "success", Method: "POST", URL: "/login", Body: `{"email":"test@test.io","password":"pass"}`, Header: nil, WantStatus: http.StatusOK, WantResponse: `{"token":"token-100"}`},
		{Name: "bad credential", Method: "POST", URL: "/login", Body: `{"email":"test@test.io","password":"wrong pass"}`, Header: nil, WantStatus: http.StatusUnauthorized, WantResponse: ""},
		{Name: "bad json", Method: "POST", URL: "/login", Body: `"email":"test@test.io","password":"wrong pass"}`, Header: nil, WantStatus: http.StatusBadRequest, WantResponse: ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
