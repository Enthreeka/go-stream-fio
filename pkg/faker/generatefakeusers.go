package faker

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"os"
	"time"
)

func GenerateUsers() {
	fakeData := generateFakeData(20000)

	err := saveDataToFile(fakeData, "fake_data.json")
	if err != nil {
		fmt.Printf("Failed to save data to file: %v\n", err)
	} else {
		fmt.Println("Data saved to fake_data.json")
	}
}

func generateFakeData(count int) Data {
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

	countryCodes := []string{"AZ", "AM", "BY", "KZ", "KG", "MD", "RU", "TJ", "TM", "UZ"}

	var fakeUsers []FakeUser
	for i := 0; i < count; i++ {
		randomlastName := gofakeit.RandomString(lastName)
		randomfirstName := gofakeit.RandomString(firstName)
		randomCountryCode := gofakeit.RandomString(countryCodes)
		if i%2 == 0 {
			fakeUser := FakeUser{
				ID:        i + 1,
				Firstname: randomfirstName,
				Lastname:  randomlastName,
				Birthday:  gofakeit.Date().Format("2006-01-02"),
				Gender:    gofakeit.Gender(),
				Address: Address{
					CountryCode: randomCountryCode,
				},
			}
			fakeUsers = append(fakeUsers, fakeUser)
		} else {
			fakeUser := FakeUser{
				ID:        i + 1,
				Firstname: gofakeit.FirstName(),
				Lastname:  gofakeit.LastName(),
				Birthday:  gofakeit.Date().Format("2006-01-02"),
				Gender:    gofakeit.Gender(),
				Address: Address{
					CountryCode: gofakeit.CountryAbr(),
				},
			}
			fakeUsers = append(fakeUsers, fakeUser)
		}
	}

	return Data{
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
