package jwt

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

// AppClaims represent the claims parsed from JWT access token.
type AppClaims struct {
	ClientID string
	Sub      string
	Scopes   []string
}

// ParseClaims parses JWT claims into AppClaims.
func (c *AppClaims) ParseClaims(claims jwt.MapClaims) error {
	clientID, ok := claims["client_id"]
	if !ok {
		return errors.New("could not parse claim id")
	}

	c.ClientID = clientID.(string)

	sub, ok := claims["sub"]
	if !ok {
		return errors.New("could not parse claim sub")
	}
	c.Sub = sub.(string)

	scp, ok := claims["scopes"]
	if !ok {
		return errors.New("could not parse claims roles")
	}

	// scopes is an array already
	var scopes []string
	if scp != nil {
		for _, v := range scp.([]interface{}) {
			scopes = append(scopes, v.(string))
		}
	}
	c.Scopes = scopes

	return nil
}
