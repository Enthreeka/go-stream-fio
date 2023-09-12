package redis

import (
	"context"
	"encoding/json"
	"github.com/NASandGAP/go-stream-fio/internal/entity"
	"github.com/NASandGAP/go-stream-fio/internal/repo"
	"github.com/NASandGAP/go-stream-fio/pkg/redis"
	"time"
)

type userRepoRedis struct {
	*redis.Redis
}

func NewUserRepoRedis(redis *redis.Redis) repo.User {
	return &userRepoRedis{
		redis,
	}
}

func (u userRepoRedis) Create(ctx context.Context, user *entity.User) error {
	bytesUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = u.Rds.Set(ctx, user.ID, bytesUser, 360*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (u userRepoRedis) GetByID(ctx context.Context, id string) (*entity.User, error) {
	user := new(entity.User)

	userBytes, err := u.Rds.Get(ctx, id).Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(userBytes, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userRepoRedis) DeleteByID(ctx context.Context, id string) error {
	err := u.Rds.Del(ctx, id).Err()
	if err != nil {
		return err
	}

	return nil
}

func (u userRepoRedis) UpdateByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
