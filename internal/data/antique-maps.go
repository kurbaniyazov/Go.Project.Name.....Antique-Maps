package data

import (
	"database/sql"
	"errors"
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
		INSERT INTO antiqueMaps (title, year, country, condition, type)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version`

	args := []interface{}{antiqueMaps.Title, antiqueMaps.Year, antiqueMaps.Country, antiqueMaps.Condition, antiqueMaps.Type}

	return a.DB.QueryRow(query, args...).Scan(&antiqueMaps.ID, &antiqueMaps.CreatedAt, &antiqueMaps.Version)
}

type AntiqueMapsModel struct {
	DB *sql.DB
}

func (a AntiqueMapsModel) Get(id int64) (*AntiqueMaps, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT id, created_at, title, year, country, condition, type, version
		FROM antiqueMaps
		WHERE id = $1`
	var antiqueMaps AntiqueMaps

	err := a.DB.QueryRow(query, id).Scan(
		&antiqueMaps.ID,
		&antiqueMaps.CreatedAt,
		&antiqueMaps.Title,
		&antiqueMaps.Year,
		&antiqueMaps.Country,
		&antiqueMaps.Condition,
		&antiqueMaps.Type,
		&antiqueMaps.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &antiqueMaps, nil
}
func (a AntiqueMapsModel) Update(antiqueMaps *AntiqueMaps) error {
	query := `
		UPDATE antiqueMaps
		SET title = $1, year = $2, country = $3, condition = $4, type = $5, version = version + 1
		WHERE id = $6 AND version = $7
		RETURNING version`

	args := []interface{}{
		antiqueMaps.Title,
		antiqueMaps.Year,
		antiqueMaps.Country,
		antiqueMaps.Condition,
		antiqueMaps.Type,
		antiqueMaps.ID,
		antiqueMaps.Version,
	}

	err := a.DB.QueryRow(query, args...).Scan(&antiqueMaps.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}
func (a AntiqueMapsModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM antiqueMaps
		WHERE id = $1`

	result, err := a.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
