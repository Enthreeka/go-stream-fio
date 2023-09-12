package server

import (
	"encoding/json"
	"github.com/NASandGAP/go-stream-fio/internal/config"
	"github.com/NASandGAP/go-stream-fio/pkg/faker"
	"github.com/NASandGAP/go-stream-fio/pkg/logger"
)

func Run(cfg *config.Config, log *logger.Logger) error {

	fakeUsers, err := faker.FakeUsers(15)
	if err != nil {
		log.Error("%v", err)
	}

	usersByte, _ := json.MarshalIndent(fakeUsers, " ", " ")
	log.Info("%s", string(usersByte))

	return nil
}
