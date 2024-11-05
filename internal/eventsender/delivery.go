package eventsender

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"sendgrid-mock/internal/model"
	"time"
)

func (s *Service) TriggerDeliveryEvent(ctx context.Context, message model.Message, err error) {
	if !s.config.Event {
		return
	}

	event := map[string]any{}

	for key, value := range message.CustomArgs {
		event[key] = value
	}

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

	event["email"] = message.To.Address
	event["timestamp"] = time.Now().Unix()
	event["event"] = eventName
	event["reason"] = "mock service: " + eventReason
	event["sg_event_id"] = message.EventID
	event["sg_message_id"] = message.MessageID
	event["smtp-id"] = message.EventID + "." + message.MessageID + "@mock"
	event["category"] = message.Categories

	s.sendEvent(ctx, event)
}

func (s *Service) sendEvent(ctx context.Context, events ...map[string]any) {
	if !s.config.Event {
		return
	}

	body, err := json.Marshal(events)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal events")

		return
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		s.config.EventEndpoint,
		bytes.NewReader(body))

	if err != nil {
		log.Error().Err(err).Msg("failed to build request")

		return
	}

	defer request.Body.Close()

	result, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to send event")

		return
	}
	defer result.Body.Close()

	log.Info().Int("status_code", result.StatusCode).Msg("webhook response")
}
