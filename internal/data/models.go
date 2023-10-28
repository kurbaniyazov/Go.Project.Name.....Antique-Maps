package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	AntiqueMaps AntiqueMapsModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		AntiqueMaps: AntiqueMapsModel{DB: db},
	}
}
