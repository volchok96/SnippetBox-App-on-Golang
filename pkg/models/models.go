package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no suitable entry was found")

type Snippet struct {
	ID      int       // Unique identifier for the snippet
	Title   string    // Title of the snippet
	Content string    // Content of the snippet
	Created time.Time // Timestamp when the snippet was created
	Expires time.Time // Timestamp when the snippet expires
}
