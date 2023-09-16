package usecase

import (
	"context"
	"github.com/Enthreeka/go-stream-fio/internal/entity/dto"
)

type User interface {
	CreateUser(ctx context.Context, fio *dto.FIO) error
}
