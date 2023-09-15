package kafka

import "github.com/segmentio/kafka-go"

func NewWriter() *kafka.Writer {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "fio-topic",
	}
	return w
}
