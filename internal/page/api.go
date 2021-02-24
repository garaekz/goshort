package page

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/<code>", res.get)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	link, err := r.service.Get(c.Request.Context(), c.Param("code"))
	if err != nil {
		status := http.StatusTemporaryRedirect
		http.Redirect(c.Response, c.Request, "/#/404", status)
		c.Abort()
		return nil
	}

	status := http.StatusTemporaryRedirect
	http.Redirect(c.Response, c.Request, link.OriginalURL, status)
	c.Abort()
	return nil
	//return c.Write(page)
}
