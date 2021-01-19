package auth

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/auth"
)

// User is the key used to store and retrieve the user identity information in routing.Context
const User = "User"

// DefaultRealm is the default realm name for HTTP authentication. It is used by HTTP authentication based on
// Basic and Bearer.
var DefaultRealm = "API"

//CustomJWT acts the same as the JWT handler from ozzo-routing but doesnt give error if user not authenticated
func CustomJWT(verificationKey string, options ...auth.JWTOptions) routing.Handler {
	var opt auth.JWTOptions
	if len(options) > 0 {
		opt = options[0]
	}
	if opt.Realm == "" {
		opt.Realm = DefaultRealm
	}
	if opt.SigningMethod == "" {
		opt.SigningMethod = "HS256"
	}
	if opt.TokenHandler == nil {
		opt.TokenHandler = auth.DefaultJWTTokenHandler
	}
	parser := &jwt.Parser{
		ValidMethods: []string{opt.SigningMethod},
	}
	return func(c *routing.Context) error {
		header := c.Request.Header.Get("Authorization")
		message := ""
		if opt.GetVerificationKey != nil {
			verificationKey = opt.GetVerificationKey(c)
		}
		if strings.HasPrefix(header, "Bearer ") {
			token, err := parser.Parse(header[7:], func(t *jwt.Token) (interface{}, error) { return []byte(verificationKey), nil })
			if err == nil && token.Valid {
				err = opt.TokenHandler(c, token)
			}
			if err == nil {
				return nil
			}
			message = err.Error()
		}

		c.Response.Header().Set("WWW-Authenticate", `Bearer realm="`+opt.Realm+`"`)
		if message != "" {
			return nil
		}
		return nil
	}
}
