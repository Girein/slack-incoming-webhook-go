# Slack Incoming Webhook for Go

## Installation
`go get github.com/Girein/slack-incoming-webhook-go`

## Setup
Add `SLACK_WEBHOOK_URL` in your `.env` file:
```
SLACK_WEBHOOK_URL=your-slack-incoming-webhook-url
```

## Usage
```
import "github.com/Girein/slack-incoming-webhook-go"

slack.PostMessage("Hello World!")
```