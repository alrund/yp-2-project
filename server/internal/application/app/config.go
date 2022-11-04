package app

import (
	"github.com/alrund/yp-2-project/server/internal/domain/port"
)

type Config struct {
	Debug        bool   `env-default:"false"`
	MigrationDir string `env-default:"migrations"`
	RunAddress   string `env:"RUN_ADDRESS" env-default:"localhost:3000"`
	DatabaseURI  string `env:"DATABASE_URI" env-default:"postgres://dev:dev@localhost:5432/dev?sslmode=disable"`
	CipherPass   string `env:"CIPHER_PASSWORD" env-default:"J53RPX6"`
	CertFile     string `env:"CERT_FILE" env-default:"local.crt"`
	KeyFile      string `env:"KEY_FILE" env-default:"local.key"`
}

func NewConfig(loader port.ConfigLoader) (*Config, error) {
	cfg := &Config{}

	flags := NewFlags()
	cfg.Debug = flags.Debug

	if err := loader.Load(cfg); err != nil {
		return nil, err
	}

	ReadFlags(flags, cfg)

	return cfg, nil
}

func ReadFlags(f *Flags, cfg *Config) {
	if f.A != NotAvailable {
		cfg.RunAddress = f.A
	}

	if f.D != NotAvailable {
		cfg.DatabaseURI = f.D
	}
}
