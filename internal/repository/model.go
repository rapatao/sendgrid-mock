package repository

import (
	"database/sql"
	"sendgrid-mock/internal/config"
)

type Service struct {
	config config.Config
	conn   *sql.DB
}
