package sendgrid

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/xeipuuv/gojsonschema"
	"sendgrid-mock/internal/repository"
	"time"
)

var (
	//go:embed json/schema.json
	schema string

	definition = gojsonschema.NewStringLoader(schema)
)

func validate(body []byte) error {
	current := gojsonschema.NewBytesLoader(body)

	result, err := gojsonschema.Validate(definition, current)
	if err != nil {
		return err
	}

	if !result.Valid() {
		return errors.New("invalid JSON")
	}

	return nil
}

func (s *Service) persist(ctx context.Context, repo *repository.Service, body []byte) (string, error) {
	var message mail.SGMailV3

	err := json.NewDecoder(bytes.NewReader(body)).Decode(&message)
	if err != nil {
		return "", err
	}

	messageID := ulid.Make().String()

	var (
		html *string
		text *string
	)

	for _, content := range message.Content {
		switch content.Type {
		case "text/html":
			html = &content.Value
		case "text/plain":
			text = &content.Value
		default:
			return "", errors.New("unsupported content type")
		}
	}

	var events []map[string]any

	for _, personalization := range message.Personalizations {
		for _, email := range personalization.To {
			eventID := ulid.Make().String()
			eventTimestamp := time.Now()

			err = save(ctx, repo, eventID, messageID, eventTimestamp, message, email, html, text)

			if s.config.Event {
				events = append(
					events,
					generateEvent(
						err,
						email.Address,
						eventID,
						messageID,
						eventTimestamp,
						categories(message.Categories, personalization.Categories),
						customArgs(message.CustomArgs, personalization.CustomArgs),
					),
				)
			}
		}
	}

	log.Info().Any("events", events).Msg("events")

	s.send(events)

	repo.TriggerCleanup()

	return messageID, err
}

func save(
	ctx context.Context,
	repo *repository.Service,
	eventID string,
	messageID string,
	eventTimestamp time.Time,
	message mail.SGMailV3,
	email *mail.Email,
	html *string,
	text *string,
) error {
	err := repo.Save(ctx, repository.Message{
		EventID:    eventID,
		MessageID:  messageID,
		ReceivedAt: eventTimestamp,
		Subject:    message.Subject,
		From: repository.Recipient{
			Name:    message.From.Name,
			Address: message.From.Address,
		},
		To: repository.Recipient{
			Name:    email.Name,
			Address: email.Address,
		},
		Content: repository.Content{
			Html: html,
			Text: text,
		},
	})
	return err
}
