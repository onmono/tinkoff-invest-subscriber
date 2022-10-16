package config

import "os"

type TinkoffInvestConfig struct {
	token     string
	isSandbox bool
}

func (t TinkoffInvestConfig) Token() string {
	return t.token
}

func (t TinkoffInvestConfig) IsSandbox() bool {
	return t.isSandbox
}

// NewConfig returns app config.
func NewConfig() TinkoffInvestConfig {
	cfg := TinkoffInvestConfig{
		token:     os.Getenv("TINKOFF_TOKEN"),
		isSandbox: true,
	}
	return cfg
}
