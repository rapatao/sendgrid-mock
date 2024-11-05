package eventsender

import (
	"github.com/rapatao/go-injector"
	"sendgrid-mock/internal/config"
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

var _ injector.Injectable = (*Service)(nil)
