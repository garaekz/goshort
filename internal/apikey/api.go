package apikey

import (
	"net/http"

	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/pkg/log"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(authHandler)
	r.Get("/apikeys/<id>", res.get)
	r.Post("/apikeys", res.create)
	r.Delete("/apikeys/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	apikeys, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(apikeys)
}

func (r resource) create(c *routing.Context) error {
	identity := auth.CurrentUser(c.Request.Context())
	userID := identity.GetID()
	apikeys, err := r.service.Create(c.Request.Context(), userID)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(apikeys, http.StatusCreated)
}

func (r resource) delete(c *routing.Context) error {
	apikeys, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(apikeys)
}
