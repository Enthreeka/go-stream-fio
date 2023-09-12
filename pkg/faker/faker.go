package faker

import (
	"encoding/json"
	"github.com/NASandGAP/go-stream-fio/internal/entity"
	"net/http"
	"strconv"
)

func FakeUsers(quantity int) (entity.Data, error) {
	quantityStr := strconv.Itoa(quantity)
	url := `https://fakerapi.it/api/v1/persons?_quantity=` + quantityStr

	resp, err := http.Get(url)
	if err != nil {
		return entity.Data{}, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	var fakeUsersResponse entity.Data

	err = decoder.Decode(&fakeUsersResponse)
	if err != nil {
		return entity.Data{}, err
	}

	return fakeUsersResponse, nil
}
