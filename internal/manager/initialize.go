package manager

import (
	"sendgrid-mock/internal/config"
	"sendgrid-mock/internal/eventsender"
	"sendgrid-mock/internal/repository"

	"github.com/rapatao/go-injector"
)

func (s *Service) Initialize(container *injector.Container) error {
	var (
		rp    repository.Service
		event eventsender.Service
		cfg   config.Config
	)

	err := container.Get(&rp)
	if err != nil {
		return err
	}

	s.repo = &rp

	err = container.Get(&event)
	if err != nil {
		return err
	}

	s.event = &event

	err = container.Get(&cfg)
	if err != nil {
		return err
	}

	s.config = &cfg

	return nil
}

var _ injector.Injectable = (*Service)(nil)
