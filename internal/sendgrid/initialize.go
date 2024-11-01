package sendgrid

import (
	"github.com/rapatao/go-injector"
	"sendgrid-mock/internal/config"
	"sendgrid-mock/internal/repository"
)

func (s *Service) Initialize(container *injector.Container) error {
	var db repository.Service

	err := container.Get(&db)
	if err != nil {
		return err
	}

	s.repo = &db

	var cfg config.Config
	err = container.Get(&cfg)
	if err != nil {
		return err
	}

	s.config = &cfg

	return nil
}
