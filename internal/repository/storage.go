package repository

import (
	"context"
	"database/sql"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rapatao/go-injector"
	"github.com/rs/zerolog/log"
	"sendgrid-mock/internal/config"
)

const (
	db = "datastore.sqlite"
)

//go:embed sql/schema.sql
var schema string

//go:embed sql/insert.sql
var insertSQL string

type Service struct {
	config  config.Config
	conn    *sql.DB
	cleanup chan bool
}

func (s *Service) Save(ctx context.Context, message Message) error {
	_, err := s.conn.ExecContext(ctx, insertSQL,
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
	)

	return err
}

func (s *Service) TriggerCleanup() {
	select {
	case s.cleanup <- true:
	default:
		log.Info().Msg("skipping cleanup since other process is running")
	}
}

var _ injector.Injectable = (*Service)(nil)
