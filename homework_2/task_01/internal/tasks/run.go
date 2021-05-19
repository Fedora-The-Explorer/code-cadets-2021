package tasks

import (
	"context"
)

func RunTasks(tasks ...Task) {
	// run each task in separate goroutine
	// wait for all tasks to finish
	//
	// when first task finishes, signal to the other goroutines that application should stop
}

type Task interface {
	Start(ctx context.Context) error
}
