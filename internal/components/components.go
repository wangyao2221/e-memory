package components

import (
	"go.uber.org/zap"

	"e-memory/configs"
)

type Components struct {
	BackendLogger *zap.Logger
	AccessLogger  *zap.Logger
}

func FromConfig(config configs.Config) (Components, error) {
	return Components{
		BackendLogger: NewBackendLogger(config),
		AccessLogger:  NewAccessLogger(config),
	}, nil
}
