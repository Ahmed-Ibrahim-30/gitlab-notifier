package events

import (
	"GitLabNotifier/notifications"
	"fmt"
)

type Dispatcher struct {
	handlers map[EventType]EventHandler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: make(map[EventType]EventHandler),
	}
}

func (d *Dispatcher) Register(
	eventType EventType,
	handler EventHandler,
) {
	d.handlers[eventType] = handler
}

func (d *Dispatcher) Dispatch(
	eventType EventType,
	payload []byte,
) (*notifications.Notification, error) {

	handler, exists := d.handlers[eventType]

	if !exists {
		return nil, fmt.Errorf(
			"no handler registered for event: %s",
			eventType,
		)
	}

	return handler.Handle(payload)
}
