package company

import (
	"context"
	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/retro-board/company-service/internal/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
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

func (c *Company) companyParts(domain string) error {
	if domain == "" {
		return bugLog.Error("domain is required")
	}

	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return bugLog.Error("domain is invalid")
	}

	c.CompanyAccount.Domain = domain
	c.CompanyAccount.Subdomain = parts[0]
	c.CompanyAccount.Name = cases.Title(language.English).String(parts[0])

	return nil
}
