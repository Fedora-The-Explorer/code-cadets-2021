package http

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"code-cadets-2021/homework_2/task_01/internal/domain/models"
)

const axilisFeedURL2 = "http://18.193.121.232/axilis-feed-2"

type AxilisOfferFeedSecond struct {
	httpClient http.Client
	updates    chan models.Odd
}

func NewAxilisOfferFeedSecond(
	httpClient http.Client,
) *AxilisOfferFeedSecond {
	return &AxilisOfferFeedSecond{
		httpClient: httpClient,
		updates:    make(chan models.Odd),
	}
}

func (a *AxilisOfferFeedSecond) Start(ctx context.Context) error {
	defer close(a.updates)
	defer log.Printf("shutting down %s", a)

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-time.After(time.Second * 3):
			response, err := a.httpClient.Get(axilisFeedURL2)
			if err != nil {
				log.Println("axilis offer feed 2, http get", err)
				continue
			}
			a.processResponseSecond(ctx, response)
		}
	}
}

func (a *AxilisOfferFeedSecond) processResponseSecond(ctx context.Context, response *http.Response) {
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("axilis offer feed 2, read all", err)
		return
	}

	offerData := string(content)

	rows := strings.Split(offerData, "\n")

	for _, row := range rows {
		offerFields := strings.Split(row, ",")
		coefficient, err := strconv.ParseFloat(offerFields[3], 64)
		if err != nil {
			log.Println("coefficient parsing", err)
			return
		}

		odd := models.Odd{
			Id:          offerFields[0],
			Name:        offerFields[1],
			Match:       offerFields[2],
			Coefficient: coefficient,
			Timestamp:   time.Now(),
		}

		select {
		case <-ctx.Done():
			return
		case a.updates <- odd:
		}
	}
}

func (a *AxilisOfferFeedSecond) String() string {
	return "axillis offer feed TWO"
}
func (a *AxilisOfferFeedSecond) GetUpdates() chan models.Odd {
	return a.updates
}
