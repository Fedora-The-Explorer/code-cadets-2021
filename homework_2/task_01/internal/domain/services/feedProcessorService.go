package services

import (
	"context"
)

type FeedProcessorService struct {
}

func NewFeedProcessorService() *FeedProcessorService {
	// it should receive "Feed" & "Queue" interfaces through constructor
	return &FeedProcessorService{}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	// initially:
	// - get updates channel from feed interface
	// - get source channel from queue interface
	//
	// repeatedly:
	// - range over updates channel
	// - multiply each odd with 2
	// - send it to source channel
	//
	// finally:
	// - when updates channel is closed, exit
	// - when exiting, close source channel
	return nil
}

// define feed interface here

// define queue interface here
