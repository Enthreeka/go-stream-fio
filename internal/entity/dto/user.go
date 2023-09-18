package dto

import "strconv"

type FioRequest struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

type IdUserRequest struct {
	ID string `json:"id"`
}

func IsNumberInFIO(fio *FioRequest) bool {
	if _, err := strconv.Atoi(fio.Name); err == nil {
		return false
	}

	if _, err := strconv.Atoi(fio.Surname); err == nil {
		return false
	}

	if _, err := strconv.Atoi(fio.Patronymic); err == nil {
		return false
	}

	return true
}

func IsRequiredField(fio *FioRequest) bool {
	if fio.Name == "" {
		return false
	}

	if fio.Surname == "" {
		return false
	}

	return true
}
