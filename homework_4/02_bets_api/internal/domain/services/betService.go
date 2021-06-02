package services

import (
	"context"
	"github.com/pkg/errors"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/api/controllers/models"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/infrastructure/sqlite"
)

type BetService struct {
	betRepository sqlite.BetRepository
}

func NewBetService(repository sqlite.BetRepository) *BetService {
	return &BetService{
		betRepository: repository,
	}
}

func (e BetService) GetBetById(ctx context.Context, id string) (models.BetResponseDto, bool, error) {
	return e.betRepository.GetBetByID(ctx, id)
}

func (e BetService) GetBetsByUser(ctx context.Context, userId string) ([]models.BetResponseDto, error) {
	return e.betRepository.GetBetByUserId(ctx, userId)
}

func (e BetService) GetBetsByStatus(ctx context.Context, status string) ([]models.BetResponseDto, error) {
	if status != "won" && status != "lost" && status != "active" {
		return []models.BetResponseDto{}, errors.New("invalid status")
	}
	return e.betRepository.GetBetsByStatus(ctx, status)
}