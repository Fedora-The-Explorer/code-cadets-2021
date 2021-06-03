package bootstrap

import (
	"github.com/streadway/amqp"
	"github.com/superbet-group/code-cadets-2021/homework_4/03_bet_acceptance_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/03_bet_acceptance_api/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework_4/03_bet_acceptance_api/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework_4/03_bet_acceptance_api/internal/api/controllers/validators"
	"github.com/superbet-group/code-cadets-2021/homework_4/03_bet_acceptance_api/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework_4/03_bet_acceptance_api/internal/infrastructure/rabbitmq"
)

func newBetValidator() *validators.BetValidator {
	return validators.NewBetValidator()
}

func newBetPublisher(publisher rabbitmq.QueuePublisher) *rabbitmq.BetPublisher {
	return rabbitmq.NewBetPublisher(
		config.Cfg.Rabbit.PublisherExchange,
		config.Cfg.Rabbit.PublisherEventUpdateQueueQueue,
		config.Cfg.Rabbit.PublisherMandatory,
		config.Cfg.Rabbit.PublisherImmediate,
		publisher,
	)
}

func newBetService(publisher services.BetPublisher) *services.BetService {
	return services.NewBetService(publisher)
}

func newController(betValidator controllers.BetValidator, betService controllers.BetService) *controllers.Controller {
	return controllers.NewController(betValidator, betService)
}

// Api bootstraps the http server.
func Api(rabbitMqChannel *amqp.Channel) *api.WebServer {
	betValidator := newBetValidator()
	betPublisher := newBetPublisher(rabbitMqChannel)
	betService := newBetService(betPublisher)
	controller := newController(betValidator, betService)

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
