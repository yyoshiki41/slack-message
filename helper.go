package slack

import "os"

func GetEnvWebhookURL() string {
	return os.Getenv("SLACK_WEBHOOK_URL")
}
