package reader

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// ErrClosed is returned once a reader has been closed.
var ErrClosed = errors.New("reader: closed")

// FileOptions configures a File reader. Position is the next byte to read;
// Generation and Identity come from the store so a restart resumes the same
// generation, or starts a new one when the file was replaced while down.
type FileOptions struct {
	Path       string
	Position   int64
	Generation int64
	Identity   string
	Policy     entry.Policy
	Sleep      func(context.Context, time.Duration) error
}

// File follows one file with a bounded batch per read.
type File struct {
	opts     FileOptions
	pol      entry.Policy
	f        *os.File
	chunk    []byte
	identity string
	gen      int64
	pos      int64
	size     int64
	drain    int64
	buf      []byte
	bufStart int64
	dropped  int
	started  bool
	missing  bool
	state    entry.State
	closed   bool
}

// NewFile prepares a reader. The file is opened on the first Read, so a
// missing path is not an error.
func NewFile(o FileOptions) (*File, error) {
	if o.Policy == (entry.Policy{}) {
		o.Policy = entry.DefaultPolicy()
	}

	if err := o.Policy.Validate(); err != nil {
		return nil, err
	}

	if o.Path == "" {
		return nil, errors.New("reader: empty path")
	}

	if o.Generation < 1 {
		o.Generation = 1
	}

	if o.Position < 0 {
		o.Position = 0
	}

	return &File{
		opts: o, pol: o.Policy, gen: o.Generation,
		pos: o.Position, identity: o.Identity, drain: -1,
		state: entry.StateIdle,
	}, nil
}
