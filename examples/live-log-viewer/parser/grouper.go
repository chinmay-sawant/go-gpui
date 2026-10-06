package parser

import (
	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// Options bounds one grouped entry.
type Options struct {
	MaxLines int
	MaxBytes int
}

// Grouper merges continuation lines, such as stack frames, into the entry
// that opened them. It emits an entry as soon as the next primary record
// arrives, and Safe reports where a reader may resume after a crash.
type Grouper struct {
	opts    Options
	gen     int64
	pending *entry.Entry
	lines   int
	bytes   int
	safeOK  bool
	safeGen int64
	safePos int64
}

// NewGrouper returns a grouper with defaults for zero options.
func NewGrouper(opts Options) *Grouper {
	if opts.MaxLines < 1 {
		opts.MaxLines = entry.DefaultPolicy().MultilineLines
	}

	if opts.MaxBytes < 1 {
		opts.MaxBytes = entry.DefaultPolicy().MultilineBytes
	}

	return &Grouper{opts: opts}
}

// Flush emits the pending entry, if any. The Ingestor calls it on shutdown
// so an unterminated final line is stored as a partial entry.
func (g *Grouper) Flush() []entry.Entry {
	if g.pending == nil {
		return nil
	}

	return []entry.Entry{g.emit()}
}

// Safe reports the generation and position just past the last emitted entry.
// A reader may commit this checkpoint; anything after it stays unread work.
func (g *Grouper) Safe() (int64, int64, bool) {
	return g.safeGen, g.safePos, g.safeOK
}

func (g *Grouper) emit() entry.Entry {
	e := *g.pending
	g.pending = nil
	g.safeGen, g.safePos, g.safeOK = e.Generation, e.Position+int64(e.Bytes), true

	return e
}
