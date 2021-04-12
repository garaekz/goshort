package user

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/garaekz/goshort/pkg/pagination"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/users/<id>", res.get)
	r.Get("/users", res.query)
	r.Post("/register", res.create)
	r.Use(authHandler)
	// the following endpoints require a valid JWT
	r.Get("/user/links", res.links)
	r.Put("/users/<id>", res.update)
	r.Delete("/users/<id>", res.delete)
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
	users, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = users
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {

	var input CreateUserRequest
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
	var input UpdateUserRequest
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

func (r resource) links(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	users, err := r.service.Links(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = users
	return c.Write(pages)
}
