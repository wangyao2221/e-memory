package server

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"e-memory/configs"
	"e-memory/internal/components"
	"e-memory/internal/routers"
)

type Server interface {
	Run()
}

type server struct {
	host   string `json:"host"`
	port   int    `json:"port"`
	engine *gin.Engine
	mux    *routers.Mux
}

func New() (Server, error) {
	config := configs.Get()
	comps, err := components.FromConfig(config)
	if err != nil {
		zap.S().Fatalf("init components failed, err: %v", err)
		return nil, nil
	}

	engine := gin.New()
	// TODO 中间件
	engine.Use(
		ginzap.Ginzap(comps.AccessLogger, time.RFC3339, true),
		ginzap.RecoveryWithZap(comps.AccessLogger, true),
	)

	s := server{
		host:   config.Server.Host,
		port:   config.Server.Port,
		engine: engine,
		mux:    routers.NewMux(engine),
	}
	s.mux.Register(comps)

	return &s, nil
}

func (s *server) Run() {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	s.engine.Run(addr)
}
