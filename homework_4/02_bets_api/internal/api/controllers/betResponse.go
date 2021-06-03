package controllers

import (
	"context"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/api/controllers/models"
)

type BetResponse interface {
	GetBetById(ctx context.Context, id string) (models.BetResponseDto, bool, error)
	GetBetsByUser(ctx context.Context, userId string) ([]models.BetResponseDto, error)
	GetBetsByStatus(ctx context.Context, status string) ([]models.BetResponseDto, error)
}
