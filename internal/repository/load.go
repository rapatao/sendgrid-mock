package repository

import (
	"context"
	_ "embed"
	"encoding/json"
	"sendgrid-mock/internal/model"
)

var (
	//go:embed sql/search.by-id.sql
	selectByIDSQL string
)

func (s *Service) Get(ctx context.Context, eventID string) (*model.Message, error) {
	row := s.conn.QueryRowContext(ctx, selectByIDSQL, eventID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var (
		message    model.Message
		customArgs string
		categories string
	)

	err := row.Scan(
		&message.EventID,
		&message.MessageID,
		&message.ReceivedAt,
		&message.Subject,
		&message.From.Name,
		&message.From.Address,
		&message.To.Name,
		&message.To.Address,
		&message.Content.Html,
		&message.Content.Text,
		&customArgs,
		&categories,
	)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(customArgs), &message.CustomArgs)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(categories), &message.Categories)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
