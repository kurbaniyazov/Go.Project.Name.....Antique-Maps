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

func (a AntiqueMapsModel) Insert(antiqueMaps *AntiqueMaps) error {

	query := `
		INSERT INTO antiquemaps (title, year, country, condition, type)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version`

	args := []interface{}{antiqueMaps.Title, antiqueMaps.Year, antiqueMaps.Country, antiqueMaps.Condition, antiqueMaps.Type}

	return a.DB.QueryRow(query, args...).Scan(&antiqueMaps.ID, &antiqueMaps.CreatedAt, &antiqueMaps.Version)
}

type AntiqueMapsModel struct {
	DB *sql.DB
}

func (a AntiqueMapsModel) Get(id int64) (*AntiqueMaps, error) {
	return nil, nil
}
func (a AntiqueMapsModel) Update(antiqueMaps *AntiqueMaps) error {
	return nil
}
func (a AntiqueMapsModel) Delete(id int64) error {
	return nil
}
