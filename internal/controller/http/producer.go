package http

import (
	"context"
	"encoding/json"
	"github.com/Enthreeka/go-stream-fio/internal/controller/tcp"
	"github.com/Enthreeka/go-stream-fio/internal/entity/dto"
	"github.com/Enthreeka/go-stream-fio/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
	"time"
)

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

func (p *producerHandler) MessageHandler(c *fiber.Ctx) error {

	fio := dto.FioRequest{}
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

	p.log.Info("send message: [%v]", fio)
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
