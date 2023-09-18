package main

import (
	"encoding/json"
	"fmt"
	"github.com/Enthreeka/go-stream-fio/internal/entity"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	fakeData := generateFakeData(20000)

	// Замените "fake_data.json" на путь и имя файла, в который вы хотите сохранить данные.
	err := saveDataToFile(fakeData, "fake_data.json")
	if err != nil {
		fmt.Printf("Failed to save data to file: %v\n", err)
	} else {
		fmt.Println("Data saved to fake_data.json")
	}
}

func generateFakeData(count int) entity.Data {
	gofakeit.Seed(time.Now().UnixNano())

	firstName := []string{
		"Nikita", "Artem", "Dmitry", "Evgeny", "Denis",
		"Ilya", "Sergey", "Vladimir", "Andrey", "Alexey",
		"Maxim", "Pavel", "Anton", "Ivan", "Alexander",
		"Konstantin", "Mikhail", "Semen", "Grigory", "Stepan",
		"Valentin", "Egor", "Nikolay", "Oleg", "Yaroslav",
		"Fedor", "Victor", "Arkady", "Timofey", "Saveliy",
		"Philip", "Gleb", "Anatoly", "Arseny", "Roman",
		"Elijah", "Mark", "George", "Daniel", "Leonid",
		"Anastasia", "Maria", "Masha", "Nadezhda", "Lubov",
		"Evgenia",
	}

	lastName := []string{
		"Smirnov", "Stepanov", "Vasilyev", "Petrov", "Ivanov",
		"Fedorov", "Sokolov", "Morozov", "Kovalenko", "Kozlov",
		"Alekseev", "Lebedev", "Grigoriev", "Titov", "Zakharov",
		"Kuznetsov", "Medvedev", "Mikhailov", "Borisov", "Sergeev",
		"Karpov", "Gorbachev", "Volkov", "Sorokin", "Zaitsev",
		"Vorobiev", "Andreev", "Novikov", "Yermakov", "Orlov",
		"Rodionov", "Komarov", "Maximov", "Pavlov", "Kazakov",
		"Denisov", "Zuev", "Fomin", "Gusev", "Tikhonov",
	}

	countryCodes := []string{
		"AZ",
		"AM",
		"BY",
		"KZ",
		"KG",
		"MD",
		"RU",
		"TJ",
		"TM",
		"UZ",
	}

	var fakeUsers []entity.FakeUser
	for i := 0; i < count; i++ {
		randomlastName := gofakeit.RandomString(lastName)
		randomfirstName := gofakeit.RandomString(firstName)
		randomCountryCode := gofakeit.RandomString(countryCodes)
		if i%2 == 0 {
			fakeUser := entity.FakeUser{
				ID:        i + 1,
				Firstname: randomfirstName,
				Lastname:  randomlastName,
				Birthday:  gofakeit.Date().Format("2006-01-02"),
				Gender:    gofakeit.Gender(),
				Address: entity.Address{
					CountryCode: randomCountryCode,
				},
			}
			fakeUsers = append(fakeUsers, fakeUser)
		} else {
			fakeUser := entity.FakeUser{
				ID:        i + 1,
				Firstname: gofakeit.FirstName(),
				Lastname:  gofakeit.LastName(),
				Birthday:  gofakeit.Date().Format("2006-01-02"),
				Gender:    gofakeit.Gender(),
				Address: entity.Address{
					CountryCode: gofakeit.CountryAbr(),
				},
			}
			fakeUsers = append(fakeUsers, fakeUser)
		}
	}

	return entity.Data{
		Total: len(fakeUsers),
		Data:  fakeUsers,
	}
}

func saveDataToFile(data interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func loadDataFromFile(filename string) (entity.Data, error) {
	var data entity.Data

	file, err := os.Open(filename)
	if err != nil {
		return data, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
