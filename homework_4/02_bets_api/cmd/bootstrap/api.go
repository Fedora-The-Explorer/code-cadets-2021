package bootstrap

import (
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/cmd/config"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/api"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/api/controllers"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/domain/services"
	"github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/infrastructure/rabbitmq"
)

func newEventUpdatePublisher(publisher rabbitmq.QueuePublisher) *rabbitmq.EventUpdatePublisher {
	return rabbitmq.NewEventUpdatePublisher(
		config.Cfg.Rabbit.PublisherExchange,
		config.Cfg.Rabbit.PublisherEventUpdateQueueQueue,
		config.Cfg.Rabbit.PublisherMandatory,
		config.Cfg.Rabbit.PublisherImmediate,
		publisher,
	)
}

func newEventService(publisher services.EventUpdatePublisher) *services.EventService {
	return services.NewEventService(publisher)
}

func newController() *controllers.Controller {
	return controllers.NewController()
}

// Api bootstraps the http server.
func Api() *api.WebServer {
	//eventUpdateValidator := newEventUpdateValidator()
	//eventUpdatePublisher := newEventUpdatePublisher(rabbitMqChannel)
	//eventService := newEventService(eventUpdatePublisher)
	controller := newController()

	return api.NewServer(config.Cfg.Api.Port, config.Cfg.Api.ReadWriteTimeoutMs, controller)
}
