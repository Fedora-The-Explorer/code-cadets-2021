package services

import (
	"context"

	"code-cadets-2021/homework_2/task_01/internal/domain/models"
)

type FeedProcessorService struct {
	feed  Feed
	queue Queue
}

func NewFeedProcessorService(feed Feed, queue Queue) *FeedProcessorService {
	// it should receive "Feed" & "Queue" interfaces through constructor
	return &FeedProcessorService{
		feed:  feed,
		queue: queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	// initially:
	// - get updates channel from feed interface
	updates := f.feed.GetUpdates()
	// - get source channel from queue interface
	source := f.queue.GetSource()

	// repeatedly:
	// - range over updates channel
	for update := range updates {
		// - multiply each odd with 2
		update.Coefficient *= 2
		// - send it to source channel
		source <- update
	}
	// finally:
	// - when updates channel is closed, exit
	// - when exiting, close source channel
	close(source)
	return nil
}

type Feed interface {
	GetUpdates() chan models.Odd
}

type Queue interface {
	GetSource() chan models.Odd
}
