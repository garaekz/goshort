package customerapi

import (
	"github.com/garaekz/goshort/internal/user"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/auth"
)

type service struct {
	repo user.Repository
}

func CustomerAPIHandler(repo user.Repository) routing.Handler {
	return auth.Bearer(func(c *routing.Context, token string) (auth.Identity, error) {
		user, err := repo.GetByApiKey(c.Request.Context(), token)
		if err != nil {
			return nil, err
		}
		return user, nil
	})
}
