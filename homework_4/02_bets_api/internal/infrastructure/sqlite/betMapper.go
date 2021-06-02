package sqlite

import (
	domainmodels "github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/api/controllers/models"
	storagemodels "github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/infrastructure/sqlite/models"
)

type BetMapper interface {
	MapDomainBetToStorageBet(domainBet domainmodels.BetResponseDto) storagemodels.Bet
	MapStorageBetToDomainBet(storageBet storagemodels.Bet) domainmodels.BetResponseDto
}
