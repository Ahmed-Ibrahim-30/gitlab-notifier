package integrations

import (
	"GitLabNotifier/apps"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Message struct {
	Text    string `json:"text"`
	Content string `json:"content"`
}

func SendMessage(message string, webhookURL string, destinationApp apps.AppType) error {
	// Check if URL is set
	if webhookURL == "" {
		return fmt.Errorf("WEBHOOK_URL is Empty")
	}

	log.Printf("📤 Try Sending notification to %s", destinationApp)

	payload := Message{Text: message}
	if destinationApp == apps.Discord {
		payload = Message{Content: message}
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON marshal error: %v", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("HTTP POST error: %v", err)
	}
	defer resp.Body.Close()

	// Log Teams response
	respBody, _ := io.ReadAll(resp.Body)
	log.Printf("📨 response status: %d", resp.StatusCode)
	log.Printf("📨 response body: %s", string(respBody))

	// Check if Teams rejected the message
	if resp.StatusCode > 300 {
		return fmt.Errorf("App returned non-200 status: %d — %s", resp.StatusCode, string(respBody))
	}

	log.Printf("📤 Sending notification to %s", destinationApp)
	return nil
}
