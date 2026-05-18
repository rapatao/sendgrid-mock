package repository

import (
	"context"
	_ "embed"
	"encoding/json"
	"sendgrid-mock/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed sql/schema.sql
	schema string

	//go:embed sql/insert.sql
	insertSQL string
)

func (s *Service) Save(ctx context.Context, message *model.Message) error {
	customArgs, err := json.Marshal(message.CustomArgs)
	if err != nil {
		return err
	}

	categories, err := json.Marshal(message.Categories)
	if err != nil {
		return err
	}

	attachments, err := json.Marshal(message.Attachments)
	if err != nil {
		return err
	}

	_, err = s.conn.ExecContext(ctx, insertSQL,
		message.EventID,
		message.MessageID,
		message.ReceivedAt,
		message.Subject,
		message.From.Name,
		message.From.Address,
		message.To.Name,
		message.To.Address,
		message.Content.Html,
		message.Content.Text,
		customArgs,
		categories,
		attachments,
	)

	return err
}
