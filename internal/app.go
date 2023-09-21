package internal

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/rohitnarayan/leaderboard/internal/config"
	"github.com/rohitnarayan/leaderboard/internal/db"
	"github.com/rohitnarayan/leaderboard/internal/logger"
)

func Init() {
	config.Init()

	writeDB, err := WriteDB(config.App)
	if err != nil {
		log.Fatalf("failed to init writeDB, err: %+v", err)
	}

	readDB, err := ReadDB(config.App)
	if err != nil {
		log.Fatalf("failed to init readDB, err: %+v", err)
	}

	logger.SetupLogger(config.App.Logger)

	log.Println(writeDB, readDB)
}

func WriteDB(cfg *config.Config) (*sqlx.DB, error) {
	return db.NewDB(cfg.Database.Postgres)
}

func ReadDB(cfg *config.Config) (*sqlx.DB, error) {
	return db.NewDB(cfg.Database.Postgres)
}
