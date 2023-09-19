package service

import (
	"context"
	"testing"

	"github.com/rohitnarayan/leaderboard/internal/config"
	"github.com/rohitnarayan/leaderboard/internal/logger"
)

func TestService(t *testing.T) {
	config.InitTestConfig()
	logger.SetupLogger(config.App.Logger)

	s := &leaderBoard{}
	s.Add(context.Background(), "123", 100)
}
