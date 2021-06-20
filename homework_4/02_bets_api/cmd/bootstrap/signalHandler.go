package bootstrap

import "github.com/superbet-group/code-cadets-2021/homework_4/02_bets_api/internal/tasks"

// SignalHandler bootstraps the signal handler.
func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
