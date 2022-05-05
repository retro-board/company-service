package config

import "fmt"

type Local struct {
	KeepLocal   bool `env:"LOCAL_ONLY" envDefault:"false"`
	Development bool `env:"DEVELOPMENT" envDefault:"false"`
	Port        int  `env:"PORT" envDefault:"3000"`

	Frontend      string `env:"FRONTEND_URL" envDefault:"retro-board.it"`
	FrontendProto string `env:"FRONTEND_PROTO" envDefault:"https"`
	JWTSecret     string `env:"JWT_SECRET" envDefault:"retro-board"`
	TokenSeed     string `env:"TOKEN_SEED" envDefault:"retro-board"`
}

func buildLocal(cfg *Config) error {
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
		if secret.Key == "jwt-secret" {
			cfg.Local.JWTSecret = secret.Value
		}

		if secret.Key == "company-token" {
			cfg.Local.TokenSeed = secret.Value
		}
	}

	return nil
}
