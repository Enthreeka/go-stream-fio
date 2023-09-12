package server

import (
	"context"
	"encoding/json"
	"github.com/NASandGAP/go-stream-fio/internal/config"
	"github.com/NASandGAP/go-stream-fio/pkg/faker"
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
	fakeUsers, err := faker.FakeUsers(15)
	if err != nil {
		log.Error("%v", err)
	}

	usersByte, _ := json.MarshalIndent(fakeUsers, " ", " ")
	log.Info("%s", string(usersByte))

	return nil
}
