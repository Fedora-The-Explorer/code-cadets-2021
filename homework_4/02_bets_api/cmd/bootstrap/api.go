package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/domain/mappers"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/infrastructure/sqlite"
)

func newBetService(betRepository sqlite.BetRepository) *services.BetService {
	return services.NewBetService(betRepository)
}

func newController(response controllers.BetResponse) *controllers.Controller {
	return controllers.NewController(response)
}


func newBetMapper() *mappers.BetMapper {
	return mappers.NewBetMapper()
}

func newBetRepository(dbExecutor sqlite.DatabaseExecutor, betMapper sqlite.BetMapper) *sqlite.BetRepository {
	return sqlite.NewBetRepository(dbExecutor, betMapper)
}

// Api bootstraps the http server.
func Api(dbExecutor sqlite.DatabaseExecutor) *api.WebServer {
	mapper := newBetMapper()
	betRepository := newBetRepository(dbExecutor, mapper)

	betService := newBetService(*betRepository)
	controller := newController(betService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
