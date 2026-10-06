// Package transfer runs one download: it opens the request, streams bytes to
// a partial file beside the destination, validates the result, and finalizes
// the file. Two transports exist: HTTP for real URLs and Fake for dummy mode.
// The package never writes outside the destination directory.
package transfer

import "errors"

// Unknown marks an expected length the server did not report.
const Unknown int64 = -1

// Errors returned to the scheduler. The scheduler wraps them with the job ID.
var (
	ErrChecksum           = errors.New("transfer: checksum mismatch")
	ErrDestinationExists  = errors.New("transfer: destination already exists")
	ErrBadRange           = errors.New("transfer: server sent an unusable range")
	ErrStall              = errors.New("transfer: no data before the stall deadline")
	ErrUnsupportedResume  = errors.New("transfer: resume cannot continue this file")
	ErrBadStatus          = errors.New("transfer: unexpected HTTP status")
	ErrPartialUnwritable  = errors.New("transfer: partial file is not writable")
	ErrDestinationMissing = errors.New("transfer: destination directory is missing")
)

// Validators are the cache validators saved with a job and replayed on
// resume. Either field may be empty.
type Validators struct {
	ETag         string
	LastModified string
}

// Empty reports whether no validator was ever saved.
func (v Validators) Empty() bool { return v.ETag == "" && v.LastModified == "" }
