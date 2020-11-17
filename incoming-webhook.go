package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
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
func PostMessage(text string) {
	body, err := json.Marshal(map[string]string{"text": text})
	if err != nil {
		fmt.Println("[Slack - Incoming Webhook - Post Message] JSON encode body got error:", err)
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("[Slack - Incoming Webhook - Post Message] Create HTTP request got error:", err)
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[Slack - Incoming Webhook - Post Message] Send HTTP request got error:", err)
	}

	response.Body.Close()
}
