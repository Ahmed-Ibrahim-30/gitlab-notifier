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

type PushHandler struct{}

func (h *PushHandler) Handle(body []byte) (*notifications.Notification, error) {

	notification := &notifications.Notification{
		EventType: string(events.EventPush),
		Title:     "New Push",
		Message:   "",
	}

	var event models.PushEvent
	if err := json.Unmarshal(body, &event); err != nil {
		log.Printf("❌ Error parsing Push event: %v", err)
		return nil, fmt.Errorf("failed to parse Push event: %w", err)
	}

	branchName := strings.TrimPrefix(event.Ref, "refs/heads/")

	commitMsg := "No commit message"
	if len(event.Commits) > 0 {
		commitMsg = strings.TrimSpace(event.Commits[0].Title)
		if commitMsg == "" {
			lines := strings.Split(strings.TrimSpace(event.Commits[0].Message), "\n")
			commitMsg = lines[0]
		}
	}

	commitCount := ""
	if len(event.Commits) > 1 {
		commitCount = fmt.Sprintf("\n📌 Total commits: *%d*", len(event.Commits))
	}

	notification.Message = fmt.Sprintf(
		"📝 *%s* committed on branch *%s*\n"+
			"💬 Commit: `%s`\n"+
			"📦 Project: _%s_%s",
		event.UserName,
		branchName,
		commitMsg,
		event.Project.Name,
		commitCount,
	)
	log.Printf("✅ Push message built: %s", notification.Message)

	return notification, nil
}
