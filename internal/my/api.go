package my

import (
	"database/sql"
	"net/http"

	"github.com/garaekz/goshort/internal/apikey"
	"github.com/garaekz/goshort/internal/auth"
	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/internal/short"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/garaekz/goshort/pkg/realip"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, apikeyService apikey.Service, shortService short.Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, apikeyService, shortService, logger}
	// the following endpoints require a valid JWT
	r.Use(authHandler)
	r.Get("/me", res.me)
	r.Get("/my/shorts", res.shorts)
	r.Post("/my/create/short", res.createMyShort)
}

type resource struct {
	service       Service
	apikeyService apikey.Service
	shortService  short.Service
	logger        log.Logger
}

func (r resource) me(c *routing.Context) error {
	me, err := r.service.GetMyUser(c.Request.Context())
	if err != nil {
		return err
	}

	keys, err := r.apikeyService.GetOwned(c.Request.Context(), me.ID)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	me.Keys = &keys

	return c.Write(me)
}

func (r resource) shorts(c *routing.Context) error {
	identity := auth.CurrentUser(c.Request.Context())
	userID := identity.GetID()
	me, err := r.shortService.GetOwned(c.Request.Context(), userID)
	if err != nil {
		return err
	}
	return c.Write(me)
}

func (r resource) createMyShort(c *routing.Context) error {
	var input short.CreateShortRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	input.IP = realip.GetRealAddr(c.Request)
	short, err := r.shortService.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}
	return c.WriteWithStatus(short, http.StatusCreated)
}
