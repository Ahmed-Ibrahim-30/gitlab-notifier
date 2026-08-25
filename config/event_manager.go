package config

import (
	"GitLabNotifier/events"
)

//responsible to enable and disable gitlab Events

type EventManager struct {
	enabledEvents map[events.EventType]struct{}
}

func NewEventManager() *EventManager {
	return &EventManager{
		enabledEvents: make(map[events.EventType]struct{}),
	}
}

func (em *EventManager) EnableEvent(eventType events.EventType) {
	em.enabledEvents[eventType] = struct{}{}
}

func (em *EventManager) DisableEvent(eventType events.EventType) {
	delete(em.enabledEvents, eventType)
}

func (em *EventManager) IsEnabled(eventType events.EventType) bool {
	_, exists := em.enabledEvents[eventType]
	return exists
}

func (em *EventManager) GetEnabledEvents() []events.EventType {
	result := make([]events.EventType, 0, len(em.enabledEvents))

	for eventType := range em.enabledEvents {
		result = append(result, eventType)
	}

	return result
}
