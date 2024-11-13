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

func (s *Service) startCleaner() {
	ticker := time.NewTicker(s.config.History)

	for threshold := range ticker.C {
		err := s.cleanup(threshold)
		if err != nil {
			log.Error().Err(err).Msg("failed to clean history")
		}
	}
}

func (s *Service) cleanup(threshold time.Time) error {
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
