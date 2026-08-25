package event_handlers

import (
	"GitLabNotifier/events"
	"GitLabNotifier/models"
	"GitLabNotifier/notifications"
	"encoding/json"
	"fmt"
	"log"
)

type ReleaseHandler struct{}

func (h *ReleaseHandler) Handle(body []byte) (*notifications.Notification, error) {
	log.Println("Release Event Handler")

	notification := &notifications.Notification{
		EventType: string(events.EventRelease),
		Title:     "New Push",
		Message:   "A new push was received.",
	}

	var event models.ReleaseEvent
	if err := json.Unmarshal(body, &event); err != nil {
		log.Printf("❌ Error parsing Release event: %v", err)
		return nil, fmt.Errorf("failed to parse Release event: %w", err)
	}

	if event.Action != "create" {
		log.Printf("⏭️ Release action '%s' ignored", event.Action)
		return nil, nil
	}

	notification.Message = fmt.Sprintf(
		"🎉 New release *%s* created in _%s_",
		event.Tag,
		event.Project.Name,
	)
	log.Printf("✅ Release message built: %s", notification.Message)

	return notification, nil
}
