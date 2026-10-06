package game

// EventKind names a change Step or Apply reports.
type EventKind uint8

// The reported changes.
const (
	EventSpawn EventKind = iota
	EventLock
	EventClear
	EventLevelUp
	EventHardDrop
	EventTopOut
	EventPause
	EventRestart
)

// Event is one reported change.
type Event struct {
	Kind  EventKind
	Lines int
	Score int
	Level int
}

// String returns the event name.
func (e Event) String() string {
	names := [...]string{
		"spawn", "lock", "clear", "level-up", "hard-drop", "top-out",
		"pause", "restart",
	}
	if int(e.Kind) < len(names) {
		return names[e.Kind]
	}

	return "unknown"
}

// valid reports whether a is a replayable action.
func (a Action) valid() bool { return a >= ActionLeft && a <= ActionStart }
