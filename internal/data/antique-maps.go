package data

import (
	"maps.alexedwards.net/internal/validator"
	"time"
)

type Maps struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Country   string    `json:"country"`
	Condition string    `json:"condition,omitempty"`
	Type      string    `json:"type"`
	Version   int32     `json:"version"`
}

func ValidateMaps(v *validator.Validator, maps *Maps) {
	v.Check(maps.Title != "", "title", "must be provided")
	v.Check(len(maps.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(maps.Year != 0, "year", "must be provided")
	v.Check(maps.Year >= 1, "year", "must be greater than 1888")

}
