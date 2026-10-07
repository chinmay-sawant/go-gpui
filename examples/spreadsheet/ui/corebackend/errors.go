package corebackend

import "errors"

// ErrNoSheet reports an unknown sheet id.
var ErrNoSheet = errors.New("corebackend: no such sheet")

// ErrRangeTooLarge reports a fetch past the adapter's bounded range.
var ErrRangeTooLarge = errors.New("corebackend: range too large")

// ErrExportTooLarge reports a CSV export past the adapter's bound.
var ErrExportTooLarge = errors.New("corebackend: export too large")
