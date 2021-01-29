package rsibot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type sendMessageReqBody struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func (rbs *RsiBotService) sendMessage(text string) error {

	reqBody := &sendMessageReqBody{
		ChatID: rbs.config.TelegramChannelName,
		Text:   text,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		"https://api.telegram.org/bot"+rbs.config.TelegramBotToken+"/"+"sendMessage",
		"application/json",
		bytes.NewBuffer(reqBytes),
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status" + resp.Status)
	}

	return err
}
