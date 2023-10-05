package data

import (
	"time"
)

type Maps struct {
	ID        int64     // Unique integer ID for the antique map
	CreatedAt time.Time // Timestamp for when the antique map is added to our database
	Title     string    // Antique map title
	Year      int32     // Antique map creation year
	Country   string    // Country of origin
	Condition string    // Condition of the antique map (e.g., "Excellent," "Good," etc.)
	Type      string    // Slice of types for the antique map
	Version   int32     // The version number starts at 1 and will be incremented each
	// time the antique map information is updated
}
