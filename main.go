package main

import (
	"go.uber.org/zap"

	"e-memory/internal/server"
)

func main() {
	s, err := server.New()
	if err != nil {
		zap.S().Fatalf("failed to new server, err: %v", err)
		return
	}
	s.Run()
}
