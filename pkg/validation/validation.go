package validation

import (
	"github.com/Enthreeka/go-stream-fio/internal/entity/dto"
	"strconv"
)

func IsNumberInFIO(fio *dto.FIO) bool {
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

func IsRequiredField(fio *dto.FIO) bool {
	if fio.Name == "" {
		return false
	}

	if fio.Surname == "" {
		return false
	}

	return true
}
