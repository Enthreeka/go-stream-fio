package main

import (
	"fmt"
	"github.com/Enthreeka/go-stream-fio/internal/controller/http"
	"github.com/Enthreeka/go-stream-fio/internal/controller/tcp"
	"github.com/Enthreeka/go-stream-fio/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log := logger.New()

	producer := tcp.NewProducerKafka(log)

	app := fiber.New()

	ph := http.NewProducerHandler(producer, log)

	app.Get("/", ph.MessageHandler)

	log.Info("Starting http kafka server: localhost:8081")

	defer producer.Close()

	if err := app.Listen(fmt.Sprintf(":8081")); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

}
