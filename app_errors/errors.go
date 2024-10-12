package app_errors

import (
	"errors"
	"net/http"
)

type statusError struct {
	error
	code int
}

type StatusError interface {
	error
	Status() int
}

func (s statusError) Unwrap() error {
	return s.error
}

func (s statusError) Status() int {
	return s.code
}

func WithStatus(err error, code int) error {
	return &statusError{
		error: err,
		code:  code,
	}
}

func GetStatus(err error) int {
	if err == nil {
		return 0
	}

	var statusErr StatusError

	if errors.As(err, &statusErr) {
		return statusErr.Status()
	}

	return http.StatusInternalServerError
}
