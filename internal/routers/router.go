package routers

import (
	"github.com/gin-gonic/gin"

	"e-memory/internal/components"
)

type Router interface {
	Register(components components.Components)
}

type Mux struct {
	engine *gin.Engine
	groups map[string]*gin.RouterGroup
}

func NewMux(engine *gin.Engine) *Mux {
	return &Mux{
		engine: engine,
		groups: map[string]*gin.RouterGroup{},
	}
}

func (m *Mux) Group(path string) *gin.RouterGroup {
	if group, ok := m.groups[path]; ok {
		return group
	}

	group := m.engine.Group(path)
	m.groups[path] = group
	return group
}

func (m *Mux) Register(components components.Components) {
	routers := []Router{
		NewMemoryRouter(m, components),
	}

	for _, router := range routers {
		router.Register(components)
	}
}
