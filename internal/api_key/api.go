package api_key

import (
	"net/http"

	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/garaekz/goshort/pkg/pagination"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Use(authHandler)
	r.Get("/api-keys/<id>", res.get)
	r.Post("/api-keys", res.create)
	r.Get("/api-keys", res.query)
	r.Put("/api-keys/<id>", res.update)
	r.Delete("/api-keys/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	user, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(user)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	api_key, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = api_key
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {

	var input CreateApiKeyRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	user, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(user, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateApiKeyRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	user, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(user)
}

func (r resource) delete(c *routing.Context) error {
	user, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(user)
}
