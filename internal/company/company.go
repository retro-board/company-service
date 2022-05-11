package company

import (
	"context"
	"github.com/retro-board/company-service/internal/config"
)

type Company struct {
	Config *config.Config
	CTX    context.Context

	CompanyAccount
}

type CompanyAccount struct {
	ID         int    `json:"id"`
	OriginalID string `json:"-"`
	Name       string `json:"name"`
	Subdomain  string `json:"subdomain"`
	Domain     string `json:"domain"`
}

func NewCompany(config *config.Config) *Company {
	return &Company{
		Config: config,
		CTX:    context.Background(),
	}
}
