package eventsender

import (
	"context"
	"sendgrid-mock/internal/model"
)

func (s *Service) TriggerOpen(ctx context.Context, message *model.Message, userAgent string, ip string) {
	if !s.config.Event {
		return
	}

	event := baseEvent(message, "open")

	event["event"] = "open"
	event["useragent"] = userAgent
	event["ip"] = ip

	s.sendEvent(ctx, event)
}
