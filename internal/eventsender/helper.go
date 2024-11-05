package eventsender

import (
	"sendgrid-mock/internal/model"
	"time"
)

func baseEvent(message *model.Message) map[string]any {
	event := map[string]any{}

	for key, value := range message.CustomArgs {
		event[key] = value
	}

	event["email"] = message.To.Address
	event["timestamp"] = time.Now().Unix()
	event["sg_event_id"] = message.EventID
	event["sg_message_id"] = message.MessageID
	event["category"] = message.Categories

	return event
}
