package manager

import (
	"github.com/rapatao/go-injector"
	"sendgrid-mock/internal/repository"
)

func (s *Service) Initialize(container *injector.Container) error {
	var rp repository.Service

	err := container.Get(&rp)
	if err != nil {
		return err
	}

	s.repo = rp

	return nil
}

var _ injector.Injectable = (*Service)(nil)
