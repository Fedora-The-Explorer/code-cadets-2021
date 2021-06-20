package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/homework_2/task_01/internal/infrastructure/http"
)

func NewAxilisOfferFeedSecond() *http.AxilisOfferFeedSecond {
	return http.NewAxilisOfferFeedSecond(stdhttp.Client{})
}
