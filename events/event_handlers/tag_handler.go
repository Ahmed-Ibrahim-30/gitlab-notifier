package event_handlers

import (
	"GitLabNotifier/events"
	"GitLabNotifier/models"
	"GitLabNotifier/notifications"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type TagHandler struct{}

func (h *TagHandler) Handle(body []byte) (*notifications.Notification, error) {
	log.Println("Tag Event Handler")

	notification := &notifications.Notification{
		EventType: string(events.EventTag),
		Title:     "New Push",
		Message:   "A new push was received.",
	}

	var event models.TagEvent
	if err := json.Unmarshal(body, &event); err != nil {
		log.Printf("❌ Error parsing Tag event: %v", err)
		return nil, fmt.Errorf("failed to parse Tag event: %w", err)
	}

	tagName := strings.TrimPrefix(event.Ref, "refs/tags/")
	notification.Message = fmt.Sprintf(
		"🏷️ *%s* created new tag *%s* in _%s_",
		event.UserName,
		tagName,
		event.Project.Name,
	)
	log.Printf("✅ Tag message built: %s", notification.Message)

	return notification, nil
}
