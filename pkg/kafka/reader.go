package kafka

import "github.com/segmentio/kafka-go"

func NewKafkaReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "my-group",
		Topic:    "fio-topic",
		MaxBytes: 10e6,
	})
}
