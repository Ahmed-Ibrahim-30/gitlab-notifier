package events

import "GitLabNotifier/notifications"

type EventHandler interface {
	Handle(payload []byte) (*notifications.Notification, error)
}
