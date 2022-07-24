package routers

import (
	"e-memory/internal/components"
	"e-memory/internal/handlers/memory"
)

type MemoryRouter struct {
	mux     *Mux
	handler memory.Handler
}

func NewMemoryRouter(mux *Mux, components components.Components) Router {
	return &MemoryRouter{
		mux:     mux,
		handler: memory.New(components.BackendLogger),
	}
}

func (r *MemoryRouter) Register(components components.Components) {
	v1 := r.mux.Group("/api/v1")
	{
		v1.GET("/memories", r.handler.List())
	}
}
