package company

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/retro-board/company-service/internal/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

func (c *Company) verifyKey(key, id string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.Config.Services.KeyService.Address, key), nil)
	if err != nil {
		bugLog.Logf("failed get service keys req: %+v", err)
		return err
	}

	req.Header.Set("Authorization", c.Config.Services.KeyService.Key)
	req.Header.Set("X-User-ID", id)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		bugLog.Logf("failed get service keys: %+v", err)
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			bugLog.Infof("failed close response body: %+v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed get service keys: %s", resp.Status)
	}

	return nil
}
