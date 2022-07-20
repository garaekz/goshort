package auth

import (
	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/pkg/log"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(rg *routing.RouteGroup, service Service, logger log.Logger) {
	rg.Post("/login", login(service, logger))
}

// login returns a handler that handles user login request.
func login(service Service, logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.Read(&req); err != nil {
			logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
			return errors.BadRequest("")
		}

		token, err := service.Login(c.Request.Context(), req.Email, req.Password)
		if err != nil {
			return err
		}
		return c.Write(struct {
			Token string `json:"token"`
		}{token})
	}
}
