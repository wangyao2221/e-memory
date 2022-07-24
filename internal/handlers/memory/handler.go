package memory

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"e-memory/internal/services/memory"
)

type Handler interface {
	i()

	// List 记忆列表
	// @Tags API.memory
	// @Router /api/memories [get]
	List() gin.HandlerFunc
}

type handler struct {
	logger  *zap.Logger
	service memory.Service
}

func New(logger *zap.Logger) Handler {
	return &handler{
		logger:  logger,
		service: memory.New(),
	}
}

func (h *handler) i() {}
