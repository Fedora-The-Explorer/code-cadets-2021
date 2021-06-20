package main

import (
	bootstrap "code-cadets-2021/homework_2/task_01/cmd/bootstrap"
	"code-cadets-2021/homework_2/task_01/internal/tasks"
)

func main() {
	signalHandler := bootstrap.NewSignalHandler()

	firstFeed := bootstrap.NewAxilisOfferFeed()
	secondFeed := bootstrap.NewAxilisOfferFeedSecond()
	mergedFeeds := bootstrap.NewFeedMerger(firstFeed, secondFeed)

	queue := bootstrap.NewOrderedQueue()
	processingService := bootstrap.FeedProcessingService(mergedFeeds, queue)

	// blocking call, start "the application"
	tasks.RunTasks(signalHandler, firstFeed, secondFeed, mergedFeeds, queue, processingService)

}
