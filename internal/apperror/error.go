package apperror

import (
	"errors"
	"fmt"
)

type appError struct {
	Err     error  `json:"-"`
	Message string `json:"message"`
}

var (
	ErrFIOFailed = NewAppError(errors.New("error_in_fio"), "required fields are written incorrectly")
)

func NewAppError(Err error, Message string) *appError {
	return &appError{
		Err:     Err,
		Message: Message,
	}
}

func (a *appError) Error() string {
	return fmt.Sprintf("%s: %s", a.Err, a.Message)
}
