package postgres

import (
	"context"
	"github.com/NASandGAP/go-stream-fio/internal/entity"
	"github.com/NASandGAP/go-stream-fio/internal/repo"
	"github.com/NASandGAP/go-stream-fio/pkg/postgres"
)

type userRepoPG struct {
	*postgres.Postgres
}

func NewUserRepoPG(postgres *postgres.Postgres) repo.User {
	return &userRepoPG{
		postgres,
	}
}

func (u userRepoPG) Create(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepoPG) GetByID(ctx context.Context, id int) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepoPG) DeleteByID(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepoPG) UpdateByID(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
