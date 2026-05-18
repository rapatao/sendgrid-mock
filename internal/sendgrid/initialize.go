package sendgrid

import (
	"sendgrid-mock/internal/config"
	"sendgrid-mock/internal/eventsender"
	"sendgrid-mock/internal/repository"

	"github.com/rapatao/go-injector"
)

func (s *Service) Initialize(container *injector.Container) error {
	var (
		db    repository.Service
		cfg   config.Config
		event eventsender.Service
	)

	err := container.Get(&db)
	if err != nil {
		return err
	}

	s.repo = &db

	err = container.Get(&cfg)
	if err != nil {
		return err
	}

	s.config = &cfg

	err = container.Get(&event)
	if err != nil {
		return err
	}

	s.event = &event

	return nil
}

var _ injector.Injectable = (*Service)(nil)
