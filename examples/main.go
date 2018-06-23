package main

import (
	"fmt"

	"github.com/yyoshiki41/slack-message"
)

func main() {
	f := slack.Field{
		Title: "title",
		Value: "this is a field.",
	}
	a := slack.Attachment{
		Fields: []slack.Field{f},
	}
	p := slack.Payload{
		Text:        "Hello, worlds.",
		Channel:     "#general",
		Attachments: []slack.Attachment{a},
	}

	err := slack.Post(slack.GetEnvWebhookURL(), p)
	if err != nil {
		fmt.Println(err)
	}
}
