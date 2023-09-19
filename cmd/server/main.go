package main

import (
	"github.com/Enthreeka/go-stream-fio/internal/config"
	"github.com/Enthreeka/go-stream-fio/internal/server"
	"github.com/Enthreeka/go-stream-fio/pkg/logger"
)

func main() {
	//path := `configs/config.json`

	log := logger.New()

	cfg, err := config.New()
	if err != nil {
		log.Error("failed to load config: %v", err)
	}

	if err := server.Run(cfg, log); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}
