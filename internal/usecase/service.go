package usecase

import (
	"context"
	"github.com/Enthreeka/go-stream-fio/internal/entity"
	"github.com/Enthreeka/go-stream-fio/internal/entity/dto"
)

type User interface {
	CreateUser(ctx context.Context, fio *dto.FioRequest) error
	FilteredUser(ctx context.Context, name string) ([]entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}
