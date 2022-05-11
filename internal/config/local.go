package config

import (
	"fmt"
	bugLog "github.com/bugfixes/go-bugfixes/logs"
)

type APIKeys struct {
	UserService    string `env:"USER_API_KEY"`
	CompanyService string `env:"COMPANY_API_KEY"`
	RetroService   string `env:"RETRO_API_KEY"`
	TimerService   string `env:"TIMER_API_KEY"`
	BillingService string `env:"BILLING_API_KEY"`
}

type Local struct {
	KeepLocal   bool `env:"LOCAL_ONLY" envDefault:"false"`
	Development bool `env:"DEVELOPMENT" envDefault:"false"`
	Port        int  `env:"PORT" envDefault:"3000"`

	Frontend      string `env:"FRONTEND_URL" envDefault:"retro-board.it"`
	FrontendProto string `env:"FRONTEND_PROTO" envDefault:"https"`
	JWTSecret     string `env:"JWT_SECRET" envDefault:"retro-board"`
	TokenSeed     string `env:"TOKEN_SEED" envDefault:"retro-board"`

	APIKeys
}

func buildLocal(cfg *Config) error {
	if err := buildLocalKeys(cfg); err != nil {
		return bugLog.Errorf("failed to build local keys: %s", err.Error())
	}

	if err := buildServiceKeys(cfg); err != nil {
		return bugLog.Errorf("failed to build service keys: %s", err.Error())
	}

	return nil
}

func buildLocalKeys(cfg *Config) error {
	vaultSecrets, err := cfg.getVaultSecrets("kv/data/retro-board/local-keys")
	if err != nil {
		return err
	}

	if vaultSecrets == nil {
		return fmt.Errorf("local keys not found in vault")
	}

	secrets, err := ParseKVSecrets(vaultSecrets)
	if err != nil {
		return err
	}

	for _, secret := range secrets {
		switch secret.Key {
		case "jwt-secret":
			cfg.Local.JWTSecret = secret.Value
			break
		case "company-token":
			cfg.Local.TokenSeed = secret.Value
			break
		}
	}

	return nil
}

func buildServiceKeys(cfg *Config) error {
	vaultSecrets, err := cfg.getVaultSecrets("kv/data/retro-board/api-keys")
	if err != nil {
		return err
	}

	if vaultSecrets == nil {
		return fmt.Errorf("api keys not found in vault")
	}

	secrets, err := ParseKVSecrets(vaultSecrets)
	if err != nil {
		return err
	}

	for _, secret := range secrets {
		switch secret.Key {
		case "retro":
			cfg.Local.APIKeys.RetroService = secret.Value
			break
		case "user":
			cfg.Local.APIKeys.UserService = secret.Value
			break
		case "company":
			cfg.Local.APIKeys.CompanyService = secret.Value
			break
		case "billing":
			cfg.Local.APIKeys.BillingService = secret.Value
			break
		case "timing":
			cfg.Local.APIKeys.TimerService = secret.Value
			break
		}
	}

	return nil
}
