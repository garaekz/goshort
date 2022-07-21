package short

import (
	"net/http"

	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/garaekz/goshort/pkg/pagination"
	"github.com/garaekz/goshort/pkg/realip"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, keyHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(keyHandler)
	r.Post("/shorts", res.create)
	r.Get("/shorts/<code>", res.get)

	r.Get("/shorts", res.query)

	r.Put("/shorts/<id>", res.update)
	// the following endpoints require a valid JWT
	r.Use(authHandler)
	r.Delete("/shorts/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	short, err := r.service.Get(c.Request.Context(), c.Param("code"))
	if err != nil {
		return err
	}
	r.service.RegisterVisit(c.Request.Context(), short.Code)
	return c.Write(short)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	shorts, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = shorts
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateShortRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	input.IP = realip.GetRealAddr(c.Request)
	short, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(short, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateShortRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	short, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(short)
}

func (r resource) delete(c *routing.Context) error {
	short, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(short)
}
