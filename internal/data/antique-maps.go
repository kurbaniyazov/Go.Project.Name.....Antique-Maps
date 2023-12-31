package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return a.DB.QueryRowContext(ctx, query, args...).Scan(&antiqueMaps.ID, &antiqueMaps.CreatedAt, &antiqueMaps.Version)
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := a.DB.QueryRowContext(ctx, query, id).Scan(
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := a.DB.QueryRowContext(ctx, query, args...).Scan(&antiqueMaps.Version)
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := a.DB.ExecContext(ctx, query, id)
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

func (a AntiqueMapsModel) GetAll(title string, country string, filters Filters) ([]*AntiqueMaps, Metadata, error) {
	query := fmt.Sprintf(`
		SELECT count(*) OVER(), id, created_at, title, year, country, condition, type, version
		FROM antiqueMaps
		WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '')
		AND (LOWER(country) = LOWER($2) OR $2 = '')
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{title, country, filters.limit(), filters.offset()}
	rows, err := a.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	antiqueMapss := []*AntiqueMaps{}

	for rows.Next() {
		var antiqueMaps AntiqueMaps
		err := rows.Scan(
			&totalRecords,
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
			return nil, Metadata{}, err
		}
		antiqueMapss = append(antiqueMapss, &antiqueMaps)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return antiqueMapss, metadata, nil
}
