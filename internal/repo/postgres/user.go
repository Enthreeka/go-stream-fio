package postgres

import (
	"context"
	"github.com/Enthreeka/go-stream-fio/internal/entity"
	"github.com/Enthreeka/go-stream-fio/internal/repo"
	"github.com/Enthreeka/go-stream-fio/pkg/postgres"
	"github.com/jackc/pgx/v5"
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
	queryPerson := `INSERT INTO person (id,firstname,lastname,age) VALUES ($1,$2,$3,$4)`
	queryGender := `INSERT INTO gender (person_id,gender,probability) VALUES ($1,$2,$3)`
	queryAddress := `INSERT INTO address (person_id,country_code,probability) VALUES ($1,$2,$3)`

	tx, err := u.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
		} else {
			tx.Commit(context.TODO())
		}
	}()

	_, err = tx.Exec(ctx, queryPerson, user.ID, user.Firstname, user.Lastname, user.Age[0].Age)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, queryGender, user.ID, user.Gender[0].Gender, user.Gender[0].Probability)
	if err != nil {
		return err
	}

	for _, address := range user.Address {
		_, err = tx.Exec(ctx, queryAddress, user.ID, address.CountryCode, address.Probability)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *userRepoPG) CreatePerson(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO person (firstname,lastname) VALUES ($1,$2)`

	_, err := u.Pool.Exec(ctx, query, user.Firstname, user.Lastname)
	return err
}

func (u *userRepoPG) GetALL(ctx context.Context) ([]entity.User, error) {
	query := `SELECT person.id,
	   person.firstname,
	   person.lastname,
	   person.age,
	   gender.gender,
	   gender.probability,
	   address.country_code,
	   address.probability
	FROM person
	JOIN  address on person.id = address.person_id
	JOIN  gender  on person.id = gender.person_id`

	rows, err := u.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	usersMap := make(map[string]entity.User)
	//users := make([]entity.User, 0)
	for rows.Next() {
		var user entity.User
		var age entity.Age
		var gender entity.Gender
		var address entity.Address

		err := rows.Scan(
			&user.ID,
			&user.Firstname,
			&user.Lastname,
			&age.Age,
			&gender.Gender,
			&gender.Probability,
			&address.CountryCode,
			&address.Probability)
		if err != nil {
			return nil, err
		}

		if data, ok := usersMap[user.ID]; ok {
			data.Address = append(data.Address, address)
		} else {
			user.Age = append(user.Age, age)
			user.Gender = append(user.Gender, gender)
			user.Address = append(user.Address, address)
			usersMap[user.ID] = user
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *userRepoPG) DeleteByID(ctx context.Context, id string) error {
	query := `DELETE FROM person WHERE id = $1`

	_, err := u.Pool.Exec(ctx, query, id)
	return err
}

func (u *userRepoPG) GetByID(ctx context.Context, id string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepoPG) UpdateByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
