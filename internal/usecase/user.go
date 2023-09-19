package usecase

import (
	"context"
	"fmt"
	"github.com/Enthreeka/go-stream-fio/internal/apperror"
	"github.com/Enthreeka/go-stream-fio/internal/entity"
	"github.com/Enthreeka/go-stream-fio/internal/entity/dto"
	"github.com/Enthreeka/go-stream-fio/internal/repo"
	"github.com/Enthreeka/go-stream-fio/pkg/faker"
	"github.com/Enthreeka/go-stream-fio/pkg/logger"
	"github.com/google/uuid"
	"time"
)

type userUsecase struct {
	userRepoPG    repo.User
	userRepoRedis repo.User

	log *logger.Logger
}

func NewUserUsecase(userRepoPG repo.User, userRepoRedis repo.User, log *logger.Logger) User {
	return &userUsecase{
		userRepoPG:    userRepoPG,
		userRepoRedis: userRepoRedis,
		log:           log,
	}
}

func (u *userUsecase) CreateUser(ctx context.Context, fio *dto.FioRequest) error {
	user, err := u.enrichmentFIO(fio)
	if err != nil {
		if err == apperror.ErrNoFoundFakeUser {
			return apperror.ErrNoFoundFakeUser
		}
		return err
	}

	user.ID = uuid.New().String()

	err = u.userRepoPG.Create(ctx, user)
	if err != nil {
		return err
	}

	err = u.userRepoRedis.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) FilteredUser(ctx context.Context, name string) ([]entity.User, error) {
	u.log.Info("starting filtered users")

	users, err := u.userRepoRedis.GetALL(ctx)
	if err != nil {
		return nil, err
	}

	filteredUsers := make([]entity.User, 0)
	for _, user := range users {
		if user.Firstname == name {
			filteredUsers = append(filteredUsers, user)
		}
	}

	u.log.Info("filtered users completed successfully")
	return filteredUsers, nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, id string) error {
	u.log.Info("start deleting a user")

	err := u.userRepoPG.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	err = u.userRepoRedis.DeleteByID(ctx, id)
	if err != nil {
		return err
	}
	u.log.Info("deleting a user with [%s] has been successfully completed", id)

	return nil
}

func (u *userUsecase) CreatePerson(ctx context.Context, user *entity.User) error {
	user.ID = uuid.New().String()

	err := u.userRepoPG.CreatePerson(ctx, user)
	if err != nil {
		return err
	}

	err = u.userRepoRedis.CreatePerson(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) enrichmentFIO(fio *dto.FioRequest) (*entity.User, error) {
	fakeUsers, err := faker.NewFaker()
	if err != nil {
		u.log.Error("failed to get fake data from API: %v", err)
		return nil, err
	}

	u.log.Info("get total - [%d] fake users", fakeUsers.Total)

	user := &entity.User{}

	var nameCount float32
	addresses := make([]entity.Address, 0)
	ages := make([]int, 0)
	genders := make([]string, 0)
	var probability float32
	for _, el := range fakeUsers.Data {
		if fio.Name == el.Firstname && fio.Surname == el.Lastname {
			nameCount++

			date, err := time.Parse("2006-01-02", el.Birthday)
			if err != nil {
				u.log.Error("parsing date [%s] error: %v", el.Birthday, err)
			}

			age := time.Now().Year() - date.Year()
			ages = append(ages, age)
			genders = append(genders, el.Gender)

			address := entity.Address{
				CountryCode: el.Address.CountryCode,
				Probability: el.Address.Probability,
			}
			addresses = append(addresses, address)
		}
	}

	if len(ages) == 0 {
		return nil, apperror.ErrNoFoundFakeUser
	}

	addressesMap := make(map[string]int)
	for _, el := range addresses {
		if data, ok := addressesMap[el.CountryCode]; !ok {
			addressesMap[el.CountryCode] = 1
		} else {
			data += 1
			addressesMap[el.CountryCode] = data
		}
	}

	gendersMap := make(map[string]int)
	for _, el := range genders {
		if data, ok := gendersMap[el]; !ok {
			gendersMap[el] = 1
		} else {
			data += 1
			gendersMap[el] = data
		}
	}

	agesMap := make(map[int]int)
	for _, el := range ages {
		if data, ok := agesMap[el]; !ok {
			agesMap[el] = 1
		} else {
			data += 1
			agesMap[el] = data
		}
	}

	addresses = []entity.Address{}
	for key, value := range addressesMap {
		var a entity.Address
		probability = float32(value) / nameCount
		probabilityStr := fmt.Sprintf("%.3f", probability)

		a.CountryCode = key
		a.Probability = probabilityStr

		addresses = append(addresses, a)
	}
	user.Address = append(user.Address, addresses...)

	for key, value := range agesMap {
		var a entity.Age
		probability = float32(value) / nameCount

		a.Age = key
		a.Probability = probability

		user.Age = append(user.Age, a)
	}

	for key, value := range gendersMap {
		var g entity.Gender
		probability = float32(value) / nameCount

		g.Gender = key
		g.Probability = probability

		user.Gender = append(user.Gender, g)
	}

	user.Firstname = fio.Name
	user.Lastname = fio.Surname

	u.highProbability(user)

	return user, nil
}

func (u *userUsecase) highProbability(user *entity.User) {
	ageProbability := entity.Age{
		Probability: user.Age[0].Probability,
		Age:         user.Age[0].Age,
	}
	////probabilityStr := fmt.Sprintf("%.2f", probability)
	for _, age := range user.Age {
		if age.Probability > ageProbability.Probability {
			ageProbability.Age = age.Age
			ageProbability.Probability = age.Probability
		}
	}
	user.Age = make([]entity.Age, 0, 1)
	user.Age = append(user.Age, ageProbability)

	genderProbability := entity.Gender{}
	if len(user.Gender) > 1 {
		if user.Gender[0].Probability > user.Gender[1].Probability {
			genderProbability.Gender = user.Gender[0].Gender
			genderProbability.Probability = user.Gender[0].Probability
		} else {
			genderProbability.Gender = user.Gender[1].Gender
			genderProbability.Probability = user.Gender[1].Probability
		}

		user.Gender = make([]entity.Gender, 0, 1)
		user.Gender = append(user.Gender, genderProbability)
	}
}
