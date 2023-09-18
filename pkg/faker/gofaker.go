package faker

import (
	"encoding/json"
	"os"
)

func NewFaker() (*Data, error) {
	file, err := os.Open("fake_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var fakeUsersResponse Data

	err = decoder.Decode(&fakeUsersResponse)
	if err != nil {
		return nil, err
	}

	return &fakeUsersResponse, nil
}
