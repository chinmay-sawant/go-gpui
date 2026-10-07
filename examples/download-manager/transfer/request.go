package transfer

import (
	"context"
	"time"
)

// Request is one transfer. Partial is the resume file; Dest is the final
// name. Expected is Unknown when the length is not known. Checksum is an
// optional "sha256:<hex>" the finished file must match. Attempt is the
// zero-based try count, which lets a fake fail only the first time.
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

// Reporter receives progress from the transfer goroutine. It must not
// block; the scheduler converts each call into a coalesced event.
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
