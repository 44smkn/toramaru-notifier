package notification

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SlackNotifier struct {
	webhookUrl string
}

type SlackMessageBody struct {
	Text string `json:"text"`
}

func NewSlackNotifier(webhookurl string) *SlackNotifier {
	return &SlackNotifier{webhookUrl: webhookurl}
}

func newSlackMessageBody(text string) *SlackMessageBody {
	return &SlackMessageBody{Text: text}
}

func (s *SlackNotifier) Notify(msg string) error {
	body, err := json.Marshal(newSlackMessageBody(msg))
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, s.webhookUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
