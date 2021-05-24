package bootstrap

import "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/calculator/internal/tasks"

func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
