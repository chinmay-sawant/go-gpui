// Package transfer runs one download: it opens the request, streams bytes to
// a partial file beside the destination, validates the result, and finalizes
// the file. Two transports exist: HTTP for real URLs and Fake for dummy mode.
// The package never writes outside the destination directory.
package transfer

import (
	"context"
	"errors"
	"time"
)

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

// Request is one transfer. Partial is the resume file; Dest is the final
// name. Expected is Unknown when the length is not known. Checksum is an
// optional "sha256:<hex>" the finished file must match.
type Request struct {
	JobID      string
	URL        string
	Dest       string
	Partial    string
	Expected   int64
	Validators Validators
	Checksum   string
	Attempt    int
}

// Progress is one observation. Total is Unknown when no length is known.
type Progress struct {
	Done  int64
	Total int64
}

// Reporter receives progress. Callers get one call per bounded chunk; a
// reporter may be called from the transfer goroutine and must not block.
type Reporter func(Progress)

// Outcome describes a finalized transfer.
type Outcome struct {
	Path       string
	Bytes      int64
	Total      int64
	Validators Validators
	Resumed    bool
}

// Transport performs one transfer. Download streams to a partial file,
// validates, and finalizes it, or returns an error and leaves a resumable
// partial behind. Download blocks until ctx is cancelled or it finishes.
type Transport interface {
	Download(ctx context.Context, req Request, report Reporter) (Outcome, error)
}

// Clock lets tests replace time. A nil Now means time.Now.
type Clock interface {
	Now() time.Time
}
