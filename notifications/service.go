package notifications

import (
	"GitLabNotifier/apps"
	"GitLabNotifier/integrations"
	"fmt"
)

type Service struct {
	appManager *apps.AppManager
}

func NewService(
	appManager *apps.AppManager,
) *Service {
	return &Service{
		appManager: appManager,
	}
}

func (s *Service) Send(
	notification *Notification,
) error {

	for _, app := range s.appManager.GetEnabledApps() {

		err := integrations.SendMessage(notification.Message, app.URL, app.Type)

		if err != nil {
			return fmt.Errorf("failed to send notification: %w", err)
		}
	}

	return nil
}
