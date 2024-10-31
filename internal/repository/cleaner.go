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

func (s *Service) cleaner() {
	for range s.cleanup {

		threshold := time.Now().Add(-s.config.History)

		exec, err := s.conn.Exec(cleanupSQL, threshold)
		if err != nil {
			log.Error().Err(err).Msg("cleanup query is failing")

			continue
		}

		affected, err := exec.RowsAffected()
		if err != nil {
			log.Error().Err(err).Msg("unable to retrieve affected rows")

			continue
		}

		log.Info().Int64("rows", affected).Msg("cleanup succeeded")
	}

	panic("cleanup should never stop")
}
