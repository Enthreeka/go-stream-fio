package repo

import (
	"context"
	"github.com/Enthreeka/go-stream-fio/internal/entity"
)

type User interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id string) (*entity.User, error)
	DeleteByID(ctx context.Context, id string) error
	UpdateByID(ctx context.Context, id string) error
}
