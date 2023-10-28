package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Movies MapsModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MapsModel{DB: db},
	}
}
