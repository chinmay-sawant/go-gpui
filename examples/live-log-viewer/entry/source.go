package entry

import "time"

// State tracks what a source is doing right now.
type State string

const (
	StateIdle    State = "idle"
	StateLive    State = "live"
	StateMissing State = "missing"
	StatePaused  State = "paused"
	StateDone    State = "done"
	StateError   State = "error"
)

// Source kinds. A file is replayable from its committed position; dummy and
// burst are generated streams whose records cannot be reread once produced.
const (
	KindFile  = "file"
	KindDummy = "dummy"
	KindBurst = "burst"
)

// Session groups sources. A dummy session exists only in the database.
type Session struct {
	ID      SessionID
	Name    string
	Kind    string
	Created time.Time
}

// Source is one followed log source and its committed checkpoint.
type Source struct {
	ID         SourceID
	Session    SessionID
	Kind       string
	Path       string
	Label      string
	Identity   string
	Generation int64
	Position   int64
	Size       int64
	Total      int64
	State      State
	Seed       int64
	Rate       int
	Lost       int64
	Updated    time.Time
}
