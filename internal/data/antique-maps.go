package data

import (
	"time"
)

type Maps struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Country   string    `json:"country"`
	Condition string    `json:"condition"`
	Type      string    `json:"type"`
	Version   int32     `json:"version"`
}
