package eventsender

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

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
