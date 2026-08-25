package handlers

import (
	"GitLabNotifier/config"
	"GitLabNotifier/events"
	"GitLabNotifier/notifications"
	"io"
	"log"
	"net/http"
	"os"
)

type WebHookHandler struct {
	eventManager        *config.EventManager
	dispatcher          *events.Dispatcher
	notificationService *notifications.Service
}

func NewWebhookHandler(eventManager *config.EventManager, dispatcher *events.Dispatcher, notificationService *notifications.Service) *WebHookHandler {
	return &WebHookHandler{
		eventManager:        eventManager,
		dispatcher:          dispatcher,
		notificationService: notificationService,
	}
}

func (h *WebHookHandler) Handle(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Gitlab-Token")
	gitlabSecretToken := os.Getenv("GITLAB_SECRET_TOKEN")

	if token != gitlabSecretToken {
		http.Error(w, "Unauthorized .. Token not match !! ", http.StatusUnauthorized)
		return
	}

	eventType := events.EventType(r.Header.Get("X-Gitlab-Event"))
	log.Printf("📥 Received event: '%s'", eventType)

	if !h.eventManager.IsEnabled(eventType) {
		log.Printf("⚠️ Disabled event received: '%s'", eventType)
		w.WriteHeader(http.StatusOK)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf("📦 Raw body: %s", string(body))

	notification, err := h.dispatcher.Dispatch(eventType, body)

	if err != nil {
		log.Printf("❌ Failed to process event '%s': %v", eventType, err)
		http.Error(w, "failed to process event", http.StatusInternalServerError)
		return
	}

	if notification == nil {
		log.Printf("⏭️ Event '%s' produced no notification", eventType)
		w.WriteHeader(http.StatusOK)
		return
	}

	// Send the generated notification to Teams.
	if err := h.notificationService.Send(notification); err != nil {
		log.Printf("❌ Failed to send notification: %v", err)
		http.Error(w, "failed to send notification", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
