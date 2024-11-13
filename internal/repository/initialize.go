package repository

import (
	"database/sql"
	"github.com/rapatao/go-injector"
	"sendgrid-mock/internal/config"
)

func (s *Service) Initialize(container *injector.Container) error {
	var cfg config.Config
	err := container.Get(&cfg)
	if err != nil {
		return err
	}

	s.config = cfg

	dbConn, err := sql.Open("sqlite3", cfg.StorageFile)
	if err != nil {
		return err
	}

	s.conn = dbConn

	_, err = s.conn.Exec(schema)
	if err != nil {

		return err
	}

	go s.startCleaner()

	return nil
}

var _ injector.Injectable = (*Service)(nil)
