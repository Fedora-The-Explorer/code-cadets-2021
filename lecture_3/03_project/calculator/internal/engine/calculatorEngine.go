package engine

import (
	"context"
	"log"
)

// Engine is the main component, responsible for consuming received bets and calculated bets,
// processing them and publishing the resulting bets.
type Calculator struct {
	consumer  Consumer
	handler   Handler
	publisher Publisher
}

// New creates and returns a new calculatorengine.
func New(consumer Consumer, handler Handler, publisher Publisher) *Calculator {
	return &Calculator{
		consumer:  consumer,
		handler:   handler,
		publisher: publisher,
	}
}

// Start will run the calculatorengine.
func (c *Calculator) Start(ctx context.Context) {
	err := c.processBetsFromController(ctx)
	if err != nil {
		log.Println("Engine failed to process bets received from controller:", err)
		return
	}

	err = c.processEventUpdates(ctx)
	if err != nil {
		log.Println("Engine failed to process event updates:", err)
		return
	}

	<-ctx.Done()
}

func (c *Calculator) processBetsFromController(ctx context.Context) error {
	consumedBetsReceived, err := c.consumer.ConsumeBetsFromController(ctx)
	if err != nil {
		return err
	}

	resultingBets := c.handler.HandleBetsFromController(ctx, consumedBetsReceived)
	c.publisher.PublishCalculatedBets(ctx, resultingBets)

	return nil
}

func (c *Calculator) processEventUpdates(ctx context.Context) error {
	consumedBetsCalculated, err := c.consumer.ConsumeEventUpdates(ctx)
	if err != nil {
		return err
	}

	resultingBets := c.handler.HandleEventUpdates(ctx, consumedBetsCalculated)
	c.publisher.PublishCalculatedBets(ctx, resultingBets)

	return nil
}
