package eventsender

import (
	"context"
	"sendgrid-mock/internal/model"
)

func (s *Service) TriggerDeliveryEvent(ctx context.Context, message *model.Message, err error) {
	if !s.config.Event {
		return
	}

	event := s.baseEvent(message)

	var (
		eventName   string
		eventReason string
	)

	if err == nil {
		eventName = "delivered"
		eventReason = "stored"
	} else {
		eventName = "dropped"
		eventReason = err.Error()
	}

	event["event"] = eventName
	event["response"] = "250 OK - mock service: " + eventReason

	s.sendEvent(ctx, event)
}
