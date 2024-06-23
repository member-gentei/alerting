package alerting

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/member-gentei/alerting/lib"
)

type MessagePublishedData struct {
	Message PubSubMessage
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

var DISCORD_WEBHOOK_URL string

func init() {
	DISCORD_WEBHOOK_URL = os.Getenv("DISCORD_WEBHOOK_URL")
	if DISCORD_WEBHOOK_URL == "" {
		log.Fatal("DISCORD_WEBHOOK_URL environment variable required")
	}
	functions.CloudEvent("HandlePubSubAlert", handlePubSubAlert)
}

func handlePubSubAlert(ctx context.Context, e cloudevents.Event) error {
	var mpd MessagePublishedData
	if err := e.DataAs(&mpd); err != nil {
		return fmt.Errorf("event.DataAs: %w", err)
	}
	var notif lib.AlertingNotification
	if err := json.Unmarshal(mpd.Message.Data, &notif); err != nil {
		return fmt.Errorf("error unmarshalling to AlertingNotification: %w", err)
	}
	return executeDiscordWebhook(notif)
}

func executeDiscordWebhook(notification lib.AlertingNotification) error {
	payload := url.Values{}
	// TODO: parse this out
	indented, err := json.MarshalIndent(notification.Incident, "", "  ")
	if err != nil {
		return fmt.Errorf("json.MarshalIndent: %w", err)
	}
	payload.Set("content", fmt.Sprintf("```json\n%s\n```", indented))
	r, err := http.PostForm(DISCORD_WEBHOOK_URL, payload)
	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}
	switch r.StatusCode {
	case 200, 204:
		return nil
	default:
		return fmt.Errorf("bad status code response: %d", r.StatusCode)
	}
}
