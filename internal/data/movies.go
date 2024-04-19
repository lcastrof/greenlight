package data

import (
	"time"
)

// Annotate the Movie struct with struct tags to control how the keys appear
// in the JSON-encoded output
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Use the - directive to omit field from output
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`    // Add the omitempty directive to hide the field from output IF it is empty
	Runtime   int32     `json:"runtime,omitempty"` // Add the omitempty directive
	Genres    []string  `json:"genres,omitempty"`  // Add the omitempty directive
	Version   int32     `json:"version"`
}
