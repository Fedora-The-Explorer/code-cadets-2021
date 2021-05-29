package bootstrap

import "github.com/superbet-group/code-cadets-2021/lecture_4/01_event_api/internal/tasks"

// SignalHandler bootstraps the signal handler.
func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
