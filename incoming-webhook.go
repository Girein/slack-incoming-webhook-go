package slack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var webhookURL string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	webhookURL = os.Getenv("SLACK_WEBHOOK_URL")
}

// PostMessage post a message to Incoming Webhook URL
func PostMessage(text string) error {
	body, err := json.Marshal(map[string]string{"text": text})
	if err != nil {
		return err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	err = response.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
