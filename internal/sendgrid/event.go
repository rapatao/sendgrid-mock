package sendgrid

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func (s *Service) send(events []map[string]any) {
	if !s.config.Event {
		return
	}

	body, err := json.Marshal(events)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal events")

		return
	}

	request, err := http.NewRequest(
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

func generateEvent(
	err error,
	email string,
	eventID string,
	messageID string,
	eventTimestamp time.Time,
	categories []string,
	customArgs map[string]string,
) map[string]any {
	event := map[string]any{}

	for key, value := range customArgs {
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

	event["email"] = email
	event["timestamp"] = eventTimestamp.Unix()
	event["event"] = eventName
	event["reason"] = "mock service: " + eventReason
	event["sg_event_id"] = eventID
	event["sg_message_id"] = messageID
	event["smtp-id"] = eventID + "." + messageID + "@mock"
	event["category"] = categories

	return event
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
