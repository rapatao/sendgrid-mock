package eventsender

import (
	"context"
	"sendgrid-mock/internal/model"
)

func (s *Service) TriggerClick(ctx context.Context, message *model.Message, userAgent string, ip string, url string) {
	if !s.config.Event {
		return
	}

	event := baseEvent(message)

	event["event"] = "click"
	event["useragent"] = userAgent
	event["ip"] = ip
	event["url"] = url

	s.sendEvent(ctx, event)
}
