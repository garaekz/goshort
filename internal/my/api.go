package my

import (
	"github.com/garaekz/goshort/pkg/log"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	// the following endpoints require a valid JWT
	r.Use(authHandler)
	r.Get("/me", res.me)
	r.Get("/my/shorts", res.shorts)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) me(c *routing.Context) error {
	me, err := r.service.GetMyUser(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(me)
}

func (r resource) shorts(c *routing.Context) error {
	me, err := r.service.GetMyShorts(c.Request.Context())
	if err != nil {
		return err
	}
	return c.Write(me)
}
