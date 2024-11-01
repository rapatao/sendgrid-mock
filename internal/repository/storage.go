package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"sendgrid-mock/internal/config"
)

const (
	db = "datastore.sqlite"
)

var (
	//go:embed sql/schema.sql
	schema string

	//go:embed sql/insert.sql
	insertSQL string
)

type Service struct {
	config config.Config
	conn   *sql.DB
}

func (s *Service) Save(ctx context.Context, message Message) error {
	customArgs, err := json.Marshal(message.CustomArgs)
	if err != nil {
		return err
	}

	categories, err := json.Marshal(message.Categories)
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
	)

	return err
}
