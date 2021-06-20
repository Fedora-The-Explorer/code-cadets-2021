package services

import (
	"code-cadets-2021/homework_2/task_01/internal/domain/models"
	"context"
	"log"
	"sync"
)

type FeedMerger struct {
	updates chan models.Odd
	feeds   []Feed
}

func NewFeedMerger(
	feeds []Feed,
) *FeedMerger {
	return &FeedMerger{
		updates: make(chan models.Odd),
		feeds:   feeds,
	}
}

func (f *FeedMerger) Start(ctx context.Context) error {
	output := f.GetUpdates()

	var wg sync.WaitGroup
	wg.Add(len(f.feeds))

	for _, feed := range f.feeds {
		go func(channels chan models.Odd, output chan models.Odd) {
			for channel := range channels {
				output <- channel
			}
			wg.Done()
		}(feed.GetUpdates(), output)
	}

	defer close(output)
	defer log.Printf("shutting down %s", f)
	wg.Wait()
	return nil
}

func (f *FeedMerger) GetUpdates() chan models.Odd {
	return f.updates
}

func (f *FeedMerger) String() string {
	return "feed merger"
}
