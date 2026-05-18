package eventsender

import (
	"sendgrid-mock/internal/config"

	"github.com/rapatao/go-injector"
)

type Service struct {
	config *config.Config
}

func (s *Service) Initialize(container *injector.Container) error {
	var cfg config.Config

	err := container.Get(&cfg)
	if err != nil {
		return err
	}

	s.config = &cfg

	return nil
}
