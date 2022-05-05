package user

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/retro-board/retro-service/internal/config"
	"golang.org/x/oauth2"
)

type User struct {
	Config      *config.Config
	Verifier    *oidc.IDTokenVerifier
	OAuthConfig *oauth2.Config
	CTX         context.Context
	State       string

	UserAccount
}

type UserAccount struct {
	ID         string   `json:"id"`
	OriginalID string   `json:"-"`
	Name       string   `json:"name"`
	Role       string   `json:"role"`
	Perms      []string `json:"perms"`

	jwt.RegisteredClaims
}

func NewUser(config *config.Config) *User {
	return &User{}
}
