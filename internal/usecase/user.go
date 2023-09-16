package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Enthreeka/go-stream-fio/internal/entity"
	"github.com/Enthreeka/go-stream-fio/internal/entity/dto"
	"github.com/Enthreeka/go-stream-fio/internal/repo"
	"github.com/Enthreeka/go-stream-fio/pkg/faker"
	"github.com/Enthreeka/go-stream-fio/pkg/logger"
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

func (u *userUsecase) CreateUser(ctx context.Context, fio *dto.FIO) error {
	user := u.enrichmentFIO(fio)

	str, _ := json.Marshal(user)

	u.log.Info("%v", string(str))

	//err := u.userRepoPG.Create(ctx, user)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (u *userUsecase) enrichmentFIO(fio *dto.FIO) *entity.User {
	// Get fake data with users from https://fakerapi.it/en
	// You can set a quantity to search a people
	fakeUsers, err := faker.FakeUsers(1000)
	if err != nil {
		u.log.Error("failed to get fake data from API: %v", err)
	}

	u.log.Info("get total - [%d] fake users", fakeUsers.Total)

	user := &entity.User{}

	var nameCount float32
	addresses := make([]entity.Address, 0)
	ages := make([]int, 0)
	genders := make([]string, 0)
	var probability float32
	for _, el := range fakeUsers.Data {
		if fio.Name == el.Firstname || fio.Surname == el.Lastname {
			nameCount++

			date, err := time.Parse("2006-01-02", el.Birthday)
			if err != nil {
				u.log.Error("parsing date [%s] error: %v", el.Birthday, err)
			}

			age := time.Now().Year() - date.Year()
			ages = append(ages, age)
			genders = append(genders, el.Gender)
			addresses = append(addresses, el.Address)
		}
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
		probabilityStr := fmt.Sprintf("%.2f", probability)

		a.CountryCode = key
		a.Probability = probabilityStr

		addresses = append(addresses, a)
	}
	user.Address = append(user.Address, addresses...)

	for key, value := range agesMap {
		var a entity.Age
		probability = float32(value) / nameCount
		probabilityStr := fmt.Sprintf("%.2f", probability)

		a.Age = key
		a.Probability = probabilityStr

		user.Age = append(user.Age, a)
	}

	for key, value := range gendersMap {
		var g entity.Gender
		probability = float32(value) / nameCount
		probabilityStr := fmt.Sprintf("%.2f", probability)

		g.Gender = key
		g.Probability = probabilityStr

		user.Gender = append(user.Gender, g)
	}

	user.Firstname = fio.Name
	user.Lastname = fio.Surname

	return user
}
