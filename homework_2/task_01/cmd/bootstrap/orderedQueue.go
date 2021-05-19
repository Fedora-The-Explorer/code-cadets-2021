package bootstrap

import "code-cadets-2021/lecture_2/05_offerfeed/internal/infrastructure/queue"

func NewOrderedQueue() *queue.OrderedQueue {
	return queue.NewOrderedQueue()
}
