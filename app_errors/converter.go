package app_errors

import (
	"errors"
	"net/http"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx"
)

const (
	duplicateKeyCode = "23505"
)

func ConvertRepoError(err error) error {
	var pgErr *pgx.PgError
	if errors.As(err, &pgErr) {
		switch pgErr {
		case pgx.ErrNoRows:
			return WithStatus(err, http.StatusNotFound)
		}
	}

	var pgConnErr *pgconn.PgError
	if errors.As(err, &pgConnErr) {
		switch pgConnErr.Code {
		case duplicateKeyCode:
			return WithStatus(err, http.StatusOK)
		}
	}

	return err
}
