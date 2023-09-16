package tcp

import (
	"context"
	"encoding/json"
	"github.com/Enthreeka/go-stream-fio/internal/apperror"
	"github.com/Enthreeka/go-stream-fio/internal/entity/dto"
	"github.com/Enthreeka/go-stream-fio/internal/usecase"
	kafkaClient "github.com/Enthreeka/go-stream-fio/pkg/kafka"
	"github.com/Enthreeka/go-stream-fio/pkg/logger"
	"github.com/Enthreeka/go-stream-fio/pkg/validation"
	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	Read(ctx context.Context) error
	Close() error
}

type consumerKafka struct {
	userUsecase usecase.User
	log         *logger.Logger

	r *kafka.Reader
}

func NewConsumerKafka(userUsecase usecase.User, log *logger.Logger) Consumer {
	return &consumerKafka{
		userUsecase: userUsecase,
		log:         log,
		r:           kafkaClient.NewKafkaReader(),
	}
}

func (c *consumerKafka) Read(ctx context.Context) error {

	for {
		msg, err := c.r.ReadMessage(ctx)
		if err != nil {
			return err
		}
		if err != nil {
			break
		}

		fio := &dto.FIO{}

		err = json.Unmarshal(msg.Value, fio)
		if err != nil {
			c.log.Error("failed to encode dto.FIO: %v", err)
		}

		if !validation.IsRequiredField(fio) {
			return apperror.ErrFIOFailed
		}

		if !validation.IsNumberInFIO(fio) {
			return apperror.ErrFIOFailed
		}

		err = c.userUsecase.CreateUser(context.Background(), fio)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *consumerKafka) Close() error {
	return c.r.Close()
}
