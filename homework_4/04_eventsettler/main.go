package main

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const eventAPI = "http://127.0.0.1:8080/event/update"
const betsAPI = "http://127.0.0.1:8081/bets?status=active"

type eventUpdateDto struct {
	Id      string `json:"id"`
	Outcome string `json:"outcome"`
}

type betDto struct {
	SelectionId string `json:"selectionId"`
}

func getActiveBets(httpClient pester.Client) ([]betDto, error) {
	httpResponse, err := httpClient.Get(betsAPI)
	if err != nil {
		return nil, err
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var dc []betDto
	error := json.Unmarshal(bodyContent, &dc)
	if error != nil {
		return nil, error
	}

	return dc, nil
}

func publishUpdates(eventUpdate eventUpdateDto) error {
	eventUpdateJson, err := json.Marshal(eventUpdate)
	if err != nil {
		return errors.WithMessage(err, "failed to marshal an event update")
	}

	_, error := http.Post(eventAPI, "application/json",
		bytes.NewBuffer(eventUpdateJson))
	if error != nil {
		return errors.WithMessage(err, "failed to post event update")
	}

	log.Printf("Sent %s", eventUpdateJson)

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	httpClient := pester.New()

	activeBets, err := getActiveBets(*httpClient)
	if err != nil {
		log.Fatalf("retrive active bets: %s", err)
	}

	matches := make(map[string]bool)
	for _, bet := range activeBets {
		matches[bet.SelectionId] = true
	}

	for key, _ := range matches {
		var outcome string
		if rand.Float64() > 0.5 {
			outcome = "lost"
		} else {
			outcome = "won"
		}

		eventUpdate := &eventUpdateDto{
			Id:      key,
			Outcome: outcome,
		}

		err := publishUpdates(*eventUpdate)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
