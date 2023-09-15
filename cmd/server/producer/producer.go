package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NASandGAP/go-stream-fio/internal/controller/tcp"
	"github.com/NASandGAP/go-stream-fio/internal/entity/dto"
	"github.com/NASandGAP/go-stream-fio/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	//kafkaConn, err := kafkaClient.New(context.Background())
	//if err != nil {
	//	log.Error("failed to connect kafka: %v", err)
	//}

	log := logger.New()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	producer := tcp.NewProducerKafka(log)

	producer.Publish(ctx)

	app := fiber.New()

	ph := NewProducerHandler(producer, log)

	app.Get("/", ph.getHandler)

	log.Info("Starting http kafka server: localhost:8080")

	defer ph.producer.Close()

	if err := app.Listen(fmt.Sprintf(":8081")); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

}

type producerHandler struct {
	producer tcp.Producer

	log *logger.Logger
}

func NewProducerHandler(producer tcp.Producer, log *logger.Logger) *producerHandler {
	return &producerHandler{
		producer: producer,
		log:      log,
	}
}

func (p *producerHandler) getHandler(c *fiber.Ctx) error {
	fio := dto.FIO{}
	err := c.BodyParser(&fio)
	if err != nil {
		p.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fioByte, err := json.Marshal(&fio)
	if err != nil {
		p.log.Error("failed to marshal: %v", err)
	}

	log.Info("%v", fio)
	err = p.producer.Publish(ctx, kafka.Message{
		Key:   []byte(fio.Name),
		Value: fioByte,
	})
	if err != nil {
		p.log.Error("failed to send msg: %v", err)
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"FIO":     fio,
		"message": "sends",
	})
}
