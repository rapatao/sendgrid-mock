package repository

import (
	"database/sql"
	"github.com/rapatao/go-injector"
	"os"
	"sendgrid-mock/internal/config"
)

func (s *Service) Initialize(container *injector.Container) error {
	file, err := os.Create(db)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	dbConn, err := sql.Open("sqlite3", db)
	if err != nil {
		return err
	}

	s.conn = dbConn

	_, err = s.conn.Exec(schema)
	if err != nil {

		return err
	}

	var cfg config.Config
	err = container.Get(&cfg)
	if err != nil {
		return err
	}

	s.config = cfg

	return nil
}

var _ injector.Injectable = (*Service)(nil)
