package integrations

import (
	"bytes"
	"consul-checks-slack/config"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func SendMessage(message interface{}) {
	timeout := time.Duration(60 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	body, err := json.Marshal(message)
	if err != nil {
		log.Println(err.Error())
	} else {
		cli, err := client.Post(*config.SlackUrl,
			"application/json", bytes.NewBuffer(body))
		defer cli.Body.Close()

		if err != nil {
			log.Println(err.Error())
		}

	}
}
