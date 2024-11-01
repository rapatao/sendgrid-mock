package sendgrid

import (
	"github.com/rs/zerolog/log"
	"time"
)

func (s *Service) startCleaner() {
	for range s.cleaner {
		threshold := time.Now().Add(-s.config.History)

		err := s.repo.Cleanup(threshold)
		if err != nil {
			log.Error().Err(err).Msg("failed to clean history")
		}
	}

	log.Fatal().Msg("cleaner processed exited")
}

func (s *Service) triggerCleaner() {
	select {
	case s.cleaner <- true:
	default:
		log.Debug().Msg("skipping cleanup since other process is running")
	}
}
