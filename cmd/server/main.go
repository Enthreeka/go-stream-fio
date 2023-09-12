package main

import (
	"github.com/NASandGAP/go-stream-fio/internal/config"
	"github.com/NASandGAP/go-stream-fio/internal/server"
	"github.com/NASandGAP/go-stream-fio/pkg/logger"
)

func main() {
	path := `configs/config.json`

	log := logger.New()

	cfg, err := config.New(path)
	if err != nil {
		log.Error("failed to load config: %v", err)
	}

	if err := server.Run(cfg, log); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}
