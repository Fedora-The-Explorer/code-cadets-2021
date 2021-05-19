package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/infrastructure/http"
)

func NewAxilisOfferFeed() *http.AxilisOfferFeed {
	return http.NewAxilisOfferFeed(stdhttp.Client{})
}
