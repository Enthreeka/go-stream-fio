package faker

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// FakeUsersAPI Get fake data with users from https://fakerapi.it/en
// You can set a quantity to search a people
func FakeUsersAPI(quantity int) (*Data, error) {
	quantityStr := strconv.Itoa(quantity)

	url := `https://fakerapi.it/api/v1/persons?_locale=ru_RU&_quantity=` + quantityStr

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var fakeUsersResponse Data

	err = decoder.Decode(&fakeUsersResponse)
	if err != nil {
		return nil, err
	}

	return &fakeUsersResponse, nil
}
