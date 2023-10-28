package data

import (
	"database/sql"
	"maps.alexedwards.net/internal/validator"
	"time"
)

type AntiqueMaps struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Country   string    `json:"country"`
	Condition string    `json:"condition,omitempty"`
	Type      string    `json:"type"`
	Version   int32     `json:"version"`
}

func ValidateAntiqueMaps(v *validator.Validator, antiqueMaps *AntiqueMaps) {
	v.Check(antiqueMaps.Title != "", "title", "must be provided")
	v.Check(len(antiqueMaps.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(antiqueMaps.Year != 0, "year", "must be provided")
	v.Check(antiqueMaps.Year >= 1, "year", "must be greater than 1888")

}

type MapsModel struct {
	DB *sql.DB
}

func (m MapsModel) Insert(antiqueMaps *AntiqueMaps) error {
	return nil
}
func (m MapsModel) Get(id int64) (*AntiqueMaps, error) {
	return nil, nil
}
func (m MapsModel) Update(antiqueMaps *AntiqueMaps) error {
	return nil
}
func (m MapsModel) Delete(id int64) error {
	return nil
}
