package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type (
	Config struct {
		Postgres   Postgres   `json:"postgres"`
		Redis      Redis      `json:"redis"`
		HTTPServer HTTPServer `json:"http_server"`
		Kafka      Kafka      `json:"kafka"`
	}

	Kafka struct {
		Topic string `json:"topic"`
	}

	Postgres struct {
		URL string `json:"url"`
	}

	Redis struct {
		Password     string `json:"password"`
		Host         string `json:"host"`
		Db           int    `json:"db"`
		MinIdleConns int    `json:"min_idle_conns"`
	}

	HTTPServer struct {
		Hostname   string `json:"hostname"`
		Port       string `json:"port"`
		TypeServer string `json:"type_server"`
	}
)

func New() (*Config, error) {
	err := godotenv.Load("configs/server.env")
	if err != nil {
		return nil, err
	}

	config := &Config{
		Postgres: Postgres{
			URL: os.Getenv("POSTGRES_URL"),
		},
		Redis: Redis{
			Password:     os.Getenv("REDIS_PASSWORD"),
			Host:         os.Getenv("REDIS_HOST"),
			Db:           parseEnvInt(os.Getenv("REDIS_DB")),
			MinIdleConns: parseEnvInt(os.Getenv("REDIS_MIN_IDLE_CONNS")),
		},
		HTTPServer: HTTPServer{
			Hostname:   os.Getenv("HTTP_HOSTNAME"),
			Port:       os.Getenv("HTTP_PORT"),
			TypeServer: os.Getenv("HTTP_TYPE_SERVER"),
		},
		Kafka: Kafka{
			Topic: os.Getenv("KAFKA_TOPIC"),
		},
	}

	return config, nil
}

func parseEnvInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}
