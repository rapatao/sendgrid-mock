package sendgrid

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"sendgrid-mock/internal/repository"
	"time"
)

func (s *Service) triggerDeliveryEvent(ctx context.Context, message repository.Message, err error) {
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
	event["category"] = categories

	s.sendEvent(ctx, event)
}

func (s *Service) sendEvent(ctx context.Context, event map[string]any) {
	if !s.config.Event {
		return
	}

	body, err := json.Marshal([]map[string]any{event})
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

func categories(values ...[]string) []string {
	var categories []string
	for _, value := range values {
		categories = append(categories, value...)
	}

	return categories
}

func customArgs(values ...map[string]string) map[string]string {
	args := map[string]string{}

	for _, value := range values {
		if value == nil {
			continue
		}

		for k, v := range value {
			args[k] = v
		}
	}

	return args
}
