package memory

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"e-memory/internal/handlers/response"
	"e-memory/internal/handlers/response/code"
)

// List 记忆列表
// @Summary 记忆列表
// @Description 记忆列表
// @Tags API.memory
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/v1/memories [get]
// @Security LoginToken
func (h *handler) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		memories, err := h.service.List()
		if err != nil {
			ctx.IndentedJSON(http.StatusBadGateway, response.Error(code.MemoryListError, err.Error()))
			return
		}
		ctx.IndentedJSON(http.StatusOK, response.Success(memories))
	}
}
