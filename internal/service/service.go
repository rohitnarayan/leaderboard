package service

import (
	"context"
	"github.com/rohitnarayan/leaderboard/internal/logger"
)

type leaderBoard struct {
}

type Leaderboard interface {
	Add(ctx context.Context, userID string, score float64) error
}

func NewLeaderBoard() Leaderboard {
	return &leaderBoard{}
}

func (s *leaderBoard) Add(ctx context.Context, userID string, score float64) error {
	logger.Infof("This message is a info message from %s", "leaderboard")
	logger.Errorf("This message is a error message from %s", "leaderboard")

	return nil
}
