package server

import (
	"context"
	"fmt"
	"github.com/NASandGAP/go-stream-fio/internal/config"
	"github.com/NASandGAP/go-stream-fio/pkg/faker"
	kafkaClient "github.com/NASandGAP/go-stream-fio/pkg/kafka"
	"github.com/NASandGAP/go-stream-fio/pkg/logger"
	"github.com/NASandGAP/go-stream-fio/pkg/postgres"
	"github.com/NASandGAP/go-stream-fio/pkg/redis"
)

func Run(cfg *config.Config, log *logger.Logger) error {

	// Connect to PostgreSQL
	psql, err := postgres.New(context.Background(), cfg.Postgres.URL)
	if err != nil {
		log.Fatal("failed to connect PostgreSQL: %v", err)
	}
	defer psql.Close()
	// Connect to Redis
	rds, err := redis.New(context.Background(), cfg)
	if err != nil {
		log.Error("redis is not working: %v", err)
	}
	defer rds.Close()

	// Get fake data with users from https://fakerapi.it/en
	// You can set a quantity to search for people
	_, err = faker.FakeUsers(15)
	if err != nil {
		log.Error("failed to get fake data from API: %v", err)
	}

	//err = w.WriteMessages(ctx, kafka.Message{
	//	Key:   []byte("Key-A"),
	//	Value: []byte("Hello World!"),
	//})
	//if err != nil {
	//	log.Error("%v", err)
	//}
	//
	//if err := w.Close(); err != nil {
	//	log.Fatal("failed to close writer:", err)
	//}
	//
	r := kafkaClient.NewKafkaReader()
	r.SetOffset(42)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

	return nil
}
