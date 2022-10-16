package app

import (
	"fmt"
	"github.com/onmono/clean-architecture/internal/config"
	"github.com/onmono/clean-architecture/pkg/database/postgres"
	"github.com/onmono/clean-architecture/pkg/logger"
)

func Run(cfg *config.Config) {
	// Logging
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use cases
	//translationUseCase := usecase.New(
	//	repo.New(pg),
	//	webapi.New(),
	//)
}
