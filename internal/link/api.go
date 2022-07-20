package link

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/garaekz/goshort/internal/errors"
	"github.com/garaekz/goshort/pkg/log"
	"github.com/garaekz/goshort/pkg/pagination"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, customerAuthHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}
	r.Use(customerAuthHandler)
	r.Post("/links", res.create)
	r.Get("/links/<code>", res.get)

	r.Use(authHandler)
	r.Get("/links", res.query)

	// the following endpoints require a valid JWT
	r.Put("/links/<id>", res.update)
	r.Delete("/links/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	link, err := r.service.GetByCode(c.Request.Context(), c.Param("code"))
	if err != nil {
		return err
	}

	return c.Write(link)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	links, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = links
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateLinkRequest
	fmt.Println(getRealAddr(c.Request))
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	link, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(link, http.StatusCreated)
}
func getRealAddr(r *http.Request) string {
	remoteIP := ""
	// the default is the originating ip. but we try to find better options because this is almost
	// never the right IP
	if parts := strings.Split(r.RemoteAddr, ":"); len(parts) == 2 {
		remoteIP = parts[0]
	}
	// If we have a forwarded-for header, take the address from there
	if xff := strings.Trim(r.Header.Get("X-Forwarded-For"), ","); len(xff) > 0 {
		addrs := strings.Split(xff, ",")
		lastFwd := addrs[len(addrs)-1]
		if ip := net.ParseIP(lastFwd); ip != nil {
			remoteIP = ip.String()
		}
		// parse X-Real-Ip header
	} else if xri := r.Header.Get("X-Real-Ip"); len(xri) > 0 {
		if ip := net.ParseIP(xri); ip != nil {
			remoteIP = ip.String()
		}
	}

	return remoteIP

}
func (r resource) update(c *routing.Context) error {
	var input UpdateLinkRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	link, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(link)
}

func (r resource) delete(c *routing.Context) error {
	link, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(link)
}
