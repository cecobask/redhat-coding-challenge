package postgres

import "fmt"

// ErrNoRows is returned when we request a row that doesn't exist
var ErrNoRows = fmt.Errorf("No matching record")
