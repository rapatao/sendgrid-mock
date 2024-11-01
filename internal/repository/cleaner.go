package repository

import (
	_ "embed"
	"github.com/rs/zerolog/log"
	"time"
)

var (
	//go:embed sql/cleanup.sql
	cleanupSQL string
)

func (s *Service) Cleanup(threshold time.Time) error {
	exec, err := s.conn.Exec(cleanupSQL, threshold)
	if err != nil {
		return err
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		log.Debug().Err(err).Msg("unable to retrieve affected rows")

		return nil
	}

	log.Debug().Int64("rows", affected).Msg("cleanup succeeded")

	return nil
}
