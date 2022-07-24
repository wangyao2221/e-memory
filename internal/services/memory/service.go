package memory

import (
	"e-memory/internal/models"
)

type Service interface {
	i()
	List() ([]models.Memory, error)
}

type service struct {
}

func New() Service {
	return &service{}
}

func (s *service) i() {}
