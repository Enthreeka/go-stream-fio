package tcp

import (
	"context"
	kafkaClient "github.com/NASandGAP/go-stream-fio/pkg/kafka"
	"github.com/NASandGAP/go-stream-fio/pkg/logger"
	"github.com/segmentio/kafka-go"
)

type Producer interface {
	Publish(ctx context.Context, msgs ...kafka.Message) error
	Close() error
}

type producerKafka struct {
	log *logger.Logger

	w *kafka.Writer
}

func NewProducerKafka(log *logger.Logger) Producer {
	return &producerKafka{
		log: log,
		w:   kafkaClient.NewWriter(),
	}
}

func (p *producerKafka) Publish(ctx context.Context, msgs ...kafka.Message) error {
	err := p.w.WriteMessages(ctx, msgs...)
	if err != nil {
		p.log.Error("failed to create message by producer: %v", err)
		return err
	}

	return nil
}

func (p *producerKafka) Close() error {
	return p.w.Close()
}
