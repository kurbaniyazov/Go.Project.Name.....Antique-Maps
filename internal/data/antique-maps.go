package data

import (
	"time"
)

type Maps struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      Year      `json:"year,omitempty"`
	Country   string    `json:"country"`
	Condition string    `json:"condition,omitempty"`
	Type      string    `json:"type"`
	Version   int32     `json:"version"`
}
