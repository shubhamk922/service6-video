package errs

import (
	"errors"
	"fmt"
)

// Data Model for error type
// errors is a package and error is an interface has one method Error

type Error struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message"`
}

// Facxtory Function
func New(code ErrCode, err error) Error {
	return Error{
		Code:    code,
		Message: err.Error(),
	}
}

func Newf(code ErrCode, format string, v ...any) Error {

	return Error{
		Code:    code,
		Message: fmt.Sprintf(format, v...),
	}

}

// Error implements the error interface
func (err Error) Error() string {
	return err.Message
}

// IsError check if the concrete type error is of type Error
func IsError(err error) bool {
	var er Error
	return errors.As(err, &er)
}

// GetError returns a copy of type Error
func GetError(err error) Error {
	var er Error
	if !errors.As(err, &er) {
		return Error{}
	}
	return er
}
