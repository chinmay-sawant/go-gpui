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
