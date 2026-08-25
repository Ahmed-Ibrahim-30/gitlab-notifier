package main

import (
	"GitLabNotifier/apps"
	"GitLabNotifier/config"
	"GitLabNotifier/events"
	"GitLabNotifier/events/event_handlers"
	"GitLabNotifier/notifications"
	"log"
	"net/http"

	"GitLabNotifier/handlers"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	eventManager := config.NewEventManager()

	eventManager.EnableEvent(events.EventPush)
	eventManager.EnableEvent(events.EventTag)
	eventManager.EnableEvent(events.EventRelease)
	eventManager.EnableEvent(events.EventMergeRequest)

	dispatcher := events.NewDispatcher()

	dispatcher.Register(
		events.EventPush,
		&event_handlers.PushHandler{},
	)

	dispatcher.Register(
		events.EventTag,
		&event_handlers.TagHandler{},
	)

	dispatcher.Register(
		events.EventRelease,
		&event_handlers.ReleaseHandler{},
	)

	dispatcher.Register(
		events.EventMergeRequest,
		&event_handlers.MergeRequestHandler{},
	)
	appManger := apps.NewAppManager()
	appManger.EnableApp(apps.Slack, "https://hooks.slack.com/services/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")
	appManger.EnableApp(apps.Discord, "https://discordapp.com/api/webhooks/1541788699556192336/FTvCc3Y57F03ykZBEpwfe_4OnS2P5hAS7kEWPW-m5hQDsjb-yKxwo5pY056ScENmZO-J")

	notificationService := notifications.NewService(appManger)

	webHookHandler := handlers.NewWebhookHandler(eventManager, dispatcher, notificationService)

	http.HandleFunc("/webhook", webHookHandler.Handle)

	log.Println("🚀 Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
