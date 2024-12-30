package products

import (
	"database/sql"
	"errors"
)

type Conf struct {
	db *sql.DB
}

func NewConf(db *sql.DB) (*Conf, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}
	return &Conf{db: db}, nil
}
