package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

var error_code = errors.New("23333")

func Dao(query string, db *sql.DB) (int, error) {
	var pk int
	err := db.QueryRow(query).Scan(&pk)

	if err == sql.ErrNoRows {
		return 0, errors.Wrap(error_code, fmt.Sprintf("queryset is none %s", query))
	}

	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("other error %s", query))
	}
	return pk, nil

}
