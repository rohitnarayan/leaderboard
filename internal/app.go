package internal

import (
	"github.com/jmoiron/sqlx"

	"github.com/rohitnarayan/leaderboard/internal/config"
)

func Init() {
	config.Init()
}

func WriteDB(cfg *config.Config) (*sqlx.DB, error) {
	return config.NewDB(cfg.Database.Postgres)
}
