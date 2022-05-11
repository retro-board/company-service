package company

import (
	"fmt"
	bugLog "github.com/bugfixes/go-bugfixes/logs"
	"github.com/jackc/pgx/v4"
)

func (c *Company) getConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(c.CTX,
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			c.Config.Database.User,
			c.Config.Database.Password,
			c.Config.Database.Host,
			c.Config.Database.Port,
			c.Config.Database.DBName))
	if err != nil {
		return nil, bugLog.Error(err)
	}

	return conn, nil
}
func (c *Company) addCompany(company *Company) error {
	conn, err := c.getConnection()
	if err != nil {
		return bugLog.Error(err)
	}
	defer func() {
		if err := conn.Close(c.CTX); err != nil {
			bugLog.Debugf("addCompany: %v", err)
		}
	}()

	var cID int
	if err := conn.QueryRow(c.CTX,
		`INSERT INTO company (name, subdomain, domain) VALUES ($1, $2, $3) RETURNING id`,
		company.CompanyAccount.Name,
		company.CompanyAccount.Subdomain,
		company.CompanyAccount.Domain).Scan(&cID); err != nil {
		return bugLog.Error(err)
	}
	c.ID = cID

	return nil
}

func (c *Company) checkCompanyExists(name string) (bool, error) {
	var exists bool

	conn, err := c.getConnection()
	if err != nil {
		return exists, bugLog.Error(err)
	}
	defer func() {
		if err := conn.Close(c.CTX); err != nil {
			bugLog.Debugf("checkCompanyExists: %v", err)
		}
	}()

	if err := conn.QueryRow(c.CTX,
		`SELECT EXISTS (SELECT 1 FROM company WHERE name = $1)`,
		name).Scan(&exists); err != nil {
		return exists, bugLog.Error(err)
	}

	return exists, nil
}
