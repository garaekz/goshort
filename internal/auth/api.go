package auth

import (
	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/pkg/log"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers registers handlers for different HTTP requests.
func RegisterHandlers(rg *routing.RouteGroup, service Service, logger log.Logger) {
	rg.Post("/login", login(service, logger))
	rg.Post("/register", register(service, logger))
	rg.Get("/verify/<id>/<token>", verify(service, logger))
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

// register returns a handler that handles user register request.
func register(service Service, logger log.Logger) routing.Handler {
	return func(c *routing.Context) error {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.Read(&req); err != nil {
			logger.With(c.Request.Context()).Errorf("invalid request: %v", err)
			return errors.BadRequest("")
		}

		err := service.Register(c.Request.Context(), req.Email, req.Password)
		if err != nil {
			return err
		}
		return c.Write(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "User registered successfully, please check your email for activation link.",
		})
	}
}

func verify(service Service, _ log.Logger) routing.Handler {
	return func(c *routing.Context) error {
		id := c.Param("id")
		token := c.Param("token")

		err := service.Verify(c.Request.Context(), id, token)
		if err != nil {
			return err
		}

		return c.Write(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "Email verified successfully, you can now login.",
		})
	}
}
