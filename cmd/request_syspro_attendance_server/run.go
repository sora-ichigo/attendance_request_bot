package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

var targetURL = "https://api.line.me/v2/bot/message/broadcast"

func Run() error {
	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		err := pushMessage()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Printf("Listening on port %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func pushMessage() error {
	channelAccessToken := os.Getenv("CHANNEL_ACCESS_TOKEN")

	body := []byte(`{
    "messages": [
      {
        "type":"text",
        "text":"シスプロの出席お願いしてもいい？"
      }
    ]
  }`)

	req, err := http.NewRequest(http.MethodPost, targetURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", channelAccessToken))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %v", res.Status)
	}

	return nil

}
