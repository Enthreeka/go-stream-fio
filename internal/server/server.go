package server

import (
	"context"
	"github.com/Enthreeka/go-stream-fio/internal/config"
	"github.com/Enthreeka/go-stream-fio/internal/controller/tcp"
	postgres2 "github.com/Enthreeka/go-stream-fio/internal/repo/postgres"
	redis2 "github.com/Enthreeka/go-stream-fio/internal/repo/redis"
	"github.com/Enthreeka/go-stream-fio/internal/usecase"
	kafkaClient "github.com/Enthreeka/go-stream-fio/pkg/kafka"
	"github.com/Enthreeka/go-stream-fio/pkg/logger"
	"github.com/Enthreeka/go-stream-fio/pkg/postgres"
	"github.com/Enthreeka/go-stream-fio/pkg/redis"
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

	conn, err := kafkaClient.New(context.Background())
	if err != nil {
		log.Fatal("failed to dial leader: %v", err)
	}

	broker := conn.Broker()
	log.Info(broker.Host, broker.ID, broker.Port, broker.Rack)

	userRepoPG := postgres2.NewUserRepoPG(psql)
	userRepoRedis := redis2.NewUserRepoRedis(rds)

	userUsecase := usecase.NewUserUsecase(userRepoPG, userRepoRedis, log)

	userConsumer := tcp.NewConsumerKafka(userUsecase, log)

	err = userConsumer.Read(context.Background())
	if err != nil {
		log.Error("%v", err)
	}

	defer userConsumer.Close()

	return nil
}
