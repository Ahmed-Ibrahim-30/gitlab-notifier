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

type MergeRequestHandler struct{}

func (h *MergeRequestHandler) Handle(body []byte) (*notifications.Notification, error) {
	log.Println("Merge Request Event Handler")

	notification := &notifications.Notification{
		EventType: string(events.EventMergeRequest),
		Title:     "New Push",
		Message:   "A new push was received.",
	}

	var event models.MergeRequestEvent
	if err := json.Unmarshal(body, &event); err != nil {
		log.Printf("❌ Error parsing MR event: %v", err)
		return nil, fmt.Errorf("failed to parse Push event: %w", err)
	}

	switch event.ObjectAttributes.Action {

	case "open", "reopen":
		reviewerNames := "No reviewer assigned"
		reviewerEmails := "No reviewer assigned"
		if len(event.Reviewers) > 0 {
			names := make([]string, len(event.Reviewers))
			emails := make([]string, len(event.Reviewers))
			for i, reviewer := range event.Reviewers {
				names[i] = reviewer.Name
				emails[i] = reviewer.Email
			}
			reviewerNames = strings.Join(names, ", ")
			reviewerEmails = strings.Join(emails, ", ")
		}

		notification.Message = fmt.Sprintf(
			"📬 *%s* opened a new Merge Request\n"+
				"🌿 Branch: *%s* → *%s*\n"+
				"👀 Reviewer: *%s*\n"+
				"👀 Reviewer Email: *%s*\n"+
				"📦 Project: _%s_",
			event.User.Name,
			event.ObjectAttributes.SourceBranch,
			event.ObjectAttributes.TargetBranch,
			reviewerNames,
			reviewerEmails,
			event.Project.Name,
		)
		log.Printf("✅ MR opened message built: %s", notification.Message)

	case "merge":
		notification.Message = fmt.Sprintf(
			"🔀 *%s* merged *%s* → *%s*\n"+
				"💬 Title: `%s`\n"+
				"📦 Project: _%s_",
			event.User.Name,
			event.ObjectAttributes.SourceBranch,
			event.ObjectAttributes.TargetBranch,
			event.ObjectAttributes.Title,
			event.Project.Name,
		)
		log.Printf("✅ MR merged message built: %s", notification.Message)

	default:
		log.Printf("⏭️ MR action '%s' ignored", event.ObjectAttributes.Action)
		return nil, nil
	}

	return notification, nil
}
