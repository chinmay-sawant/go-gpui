package storage

import (
	"errors"
)

// Sentinel errors callers switch on.
var (
	ErrClosed      = errors.New("storage: store is closed")
	ErrSchemaNewer = errors.New("storage: database schema is newer than this build; refusing to modify it")
	ErrConflict    = errors.New("storage: workbook changed since it was loaded")
	ErrNotFound    = errors.New("storage: not found")
)
