package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Token holds refresh jwt information.
type Token struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	AccountID int       `json:"-"`

	Token      string    `json:"-"`
	Expiry     time.Time `json:"-"`
	Mobile     bool      `sql:",notnull" json:"mobile"`
	Identifier string    `json:"identifier,omitempty"`
}

// Claims returns the token claims to be signed
func (t *Token) Claims() jwt.MapClaims {
	return jwt.MapClaims{
		"id":    t.ID,
		"token": t.Token,
	}
}
