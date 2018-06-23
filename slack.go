package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Attachment struct {
	AttachmentType string  `json:"attachment_type"`
	AuthorName     string  `json:"author_name"`
	AuthorLink     string  `json:"author_link"`
	AuthorIcon     string  `json:"author_icon"`
	Color          string  `json:"color,omitempty"`
	Fallback       string  `json:"fallback"`
	Fields         []Field `json:"fields"`
	Footer         string  `json:"footer"`
	FooterIcon     string  `json:"footer_icon"`
	ImageURL       string  `json:"image_url"`
	ThumbURL       string  `json:"thumb_url"`
	Pretext        string  `json:"pretext,omitempty"`
	Text           string  `json:"text"`
	Ts             uint64  `json:"ts"`
	Title          string  `json:"title"`
	TitleLink      string  `json:"title_link"`
}

type Payload struct {
	Text        string       `json:"text"`
	Channel     string       `json:"channel"`
	UserName    string       `json:"username"`
	IconEmoji   string       `json:"icon_emoji"`
	Attachments []Attachment `json:"attachments"`
}

func Post(webhook string, payload Payload) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	payloadBody := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(http.MethodPost, webhook, payloadBody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("slack-message: HTTP StatusCode is %d", resp.StatusCode)
	}
	return nil
}
