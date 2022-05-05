package company

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/retro-board/company-service/internal/config"
	"golang.org/x/oauth2"
)

type Company struct {
	Config      *config.Config
	Verifier    *oidc.IDTokenVerifier
	OAuthConfig *oauth2.Config
	CTX         context.Context
	State       string

	CompanyAccount
}

type CompanyAccount struct {
	ID         string `json:"id"`
	OriginalID string `json:"-"`
	Name       string `json:"name"`

	jwt.RegisteredClaims
}

func NewCompany(config *config.Config) *Company {
	return &Company{}
}
