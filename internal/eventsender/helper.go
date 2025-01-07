package eventsender

import (
	"sendgrid-mock/internal/model"
	"time"
)

func baseEvent(message *model.Message, status string) map[string]any {
	event := map[string]any{}

	for key, value := range message.CustomArgs {
		event[key] = value
	}

	event["email"] = message.To.Address
	event["timestamp"] = time.Now().Unix() + getExtraTimeByStatus(status)
	event["smtp-id"] = message.EventID + "." + message.MessageID + "@mock"
	event["sg_event_id"] = message.EventID
	event["sg_message_id"] = message.MessageID
	event["category"] = message.Categories

	return event
}

func getExtraTimeByStatus(status string) int64 {
	switch status {
	case "delivery":
		return 30
	case "open":
		return 45
	case "click":
		return 60
	default:
		return 0
	}
}
