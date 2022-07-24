package memory

import (
	"e-memory/internal/models"
)

func (s *service) List() ([]models.Memory, error) {
	return []models.Memory{
		{
			"1",
			"first memory",
			"first memory content",
		},
	}, nil
}
