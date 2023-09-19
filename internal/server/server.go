package server

import (
	"context"
	"fmt"
	"github.com/Enthreeka/go-stream-fio/internal/config"
	"github.com/Enthreeka/go-stream-fio/internal/controller/http"
	"github.com/Enthreeka/go-stream-fio/internal/controller/tcp"
	postgres2 "github.com/Enthreeka/go-stream-fio/internal/repo/postgres"
	redis2 "github.com/Enthreeka/go-stream-fio/internal/repo/redis"
	"github.com/Enthreeka/go-stream-fio/internal/usecase"
	kafkaClient "github.com/Enthreeka/go-stream-fio/pkg/kafka"
	"github.com/Enthreeka/go-stream-fio/pkg/logger"
	"github.com/Enthreeka/go-stream-fio/pkg/postgres"
	"github.com/Enthreeka/go-stream-fio/pkg/redis"
	"github.com/gofiber/fiber/v2"
	"sync"
	"time"
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

	brokers, err := conn.Brokers()
	if err != nil {
		return err
	}
	log.Info("kafka connected to brokers: %+v", brokers)

	userRepoPG := postgres2.NewUserRepoPG(psql)
	userRepoRedis := redis2.NewUserRepoRedis(rds)

	//_, err = userRepoPG.GetALL(context.Background())

	userUsecase := usecase.NewUserUsecase(userRepoPG, userRepoRedis, log)

	userConsumer := tcp.NewConsumerKafka(userUsecase, log)

	userHandler := http.NewUserHandler(userUsecase, log)

	app := fiber.New()

	app.Get("/users", userHandler.UserHandler)
	app.Delete("/:id", userHandler.DeleteUserHandler)
	app.Post("/user", userHandler.CreatePersonHandler)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			err = userConsumer.Read(context.Background())
			if err != nil {
				log.Error("%v", err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info("Starting http server: %s:%s", cfg.HTTPServer.TypeServer, cfg.HTTPServer.Port)

		if err = app.Listen(fmt.Sprintf(":%s", cfg.HTTPServer.Port)); err != nil {
			log.Fatal("Server listening failed:%s", err)
		}
	}()

	wg.Wait()

	defer userConsumer.Close()

	return nil
}
