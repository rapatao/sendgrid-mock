package sendgrid

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"github.com/oklog/ulid/v2"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"sendgrid-mock/internal/repository"
	"time"
)

func (s *Service) persist(ctx context.Context, body []byte) (string, error) {
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

	for _, personalization := range message.Personalizations {
		for _, email := range personalization.To {
			message := repository.Message{
				EventID:    ulid.Make().String(),
				MessageID:  messageID,
				ReceivedAt: time.Now(),
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
				CustomArgs: customArgs(message.CustomArgs, personalization.CustomArgs),
				Categories: categories(message.Categories, personalization.Categories),
			}

			err = s.repo.Save(ctx, message)
			s.triggerDeliveryEvent(ctx, message, err)
		}
	}

	s.triggerCleaner()

	return messageID, err
}
