package manager

import (
	"github.com/rapatao/go-injector"
	"sendgrid-mock/internal/eventsender"
	"sendgrid-mock/internal/repository"
)

func (s *Service) Initialize(container *injector.Container) error {
	var (
		rp    repository.Service
		event eventsender.Service
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

	return nil
}

var _ injector.Injectable = (*Service)(nil)
