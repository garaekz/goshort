package healthcheck

import (
	"net/http"
	"testing"

	"github.com/garaekz/goshort/internal/test"
	"github.com/garaekz/goshort/pkg/log"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	RegisterHandlers(router, "0.9.0")
	test.Endpoint(t, router, test.APITestCase{
		"ok", "GET", "/healthcheck", "", nil, http.StatusOK, `"OK 0.9.0"`,
	})
}
