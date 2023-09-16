package postgres

import (
	"context"
	"github.com/Enthreeka/go-stream-fio/internal/entity"
	"github.com/Enthreeka/go-stream-fio/internal/repo"
	"github.com/Enthreeka/go-stream-fio/pkg/postgres"
)

type userRepoPG struct {
	*postgres.Postgres
}

func NewUserRepoPG(postgres *postgres.Postgres) repo.User {
	return &userRepoPG{
		postgres,
	}
}

func (u *userRepoPG) Create(ctx context.Context, user *entity.User) error {
	//queryPerson := `INSERT INTO person (firstname,lastname,age) VALUES ($1,$2,$3) RETURNING id`
	//queryGender := `INSERT INTO gender (person_id,gender,probability) VALUES ($1,$2,$3)`
	//queryAddress := `INSERT INTO address (person_id,country_code,probability) VALUES ($1,$2,$3)`
	//
	//tx, err := u.Pool.BeginTx(ctx, pgx.TxOptions{})
	//if err != nil {
	//	return err
	//}
	//
	//defer func() {
	//	if err != nil {
	//		tx.Rollback(context.TODO())
	//	} else {
	//		tx.Commit(context.TODO())
	//	}
	//}()

	//var personID int
	//err = tx.QueryRow(ctx, queryPerson, user.Firstname, user.Lastname, user.Birthday).Scan(&personID)
	//if err != nil {
	//	return err
	//}
	//
	//_, err = tx.Exec(ctx, queryGender, personID, user.Gender, user.Probability)
	//if err != nil {
	//	return err
	//}
	//
	//_, err = tx.Exec(ctx, queryAddress, personID, user.Address.CountryCode, user.Address.Probability)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (u *userRepoPG) GetByID(ctx context.Context, id string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepoPG) DeleteByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepoPG) UpdateByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
