package ui

import (
	"sync/atomic"
	"time"
)

// graphHistory bounds every graph buffer, so memory and paint work never
// grow with uptime whatever rate the source delivers.
const graphHistory = 120

// panelOrder is the overview order.
var panelOrder = []string{"cpu", "mem", "disk", "net"}

// panelNames are the card titles.
var panelNames = map[string]string{
	"cpu":  "CPU",
	"mem":  "Memory",
	"disk": "Disk",
	"net":  "Network",
}

// panel is one overview card's live state.
type panel struct {
	id    string
	name  string
	graph *Graph
	last  Reading
	have  bool
}

// selection is the process detail state. id is the identity the user picked;
// found tracks whether the frozen snapshot still holds it.
type selection struct {
	id    string
	proc  Process
	found bool
	trk   Tracked
}

// state is everything the UI loop owns. Handlers and Tick are the only
// writers, so it needs no lock. gen is atomic because the poll goroutines
// read it to tag their results.
type state struct {
	nav       string
	dark      bool
	live      bool
	mode      string
	at        time.Time
	notice    string
	problems  []string
	panels    map[string]*panel
	table     *table
	sel       selection
	gen       atomic.Uint64
	h         handles
	lastGen   uint64
	haveAny   bool
	dirtyText bool
}

// newState builds the loop state for a mode.
func newState(dark, live bool) *state {
	s := &state{
		nav:    "overview",
		dark:   dark,
		live:   live,
		mode:   modeName(live),
		panels: make(map[string]*panel, len(panelOrder)),
		table:  newTable(),
	}

	for _, id := range panelOrder {
		s.panels[id] = &panel{id: id, name: panelNames[id], graph: NewGraph(graphHistory)}
	}

	return s
}

// modeName labels a mode for the UI.
func modeName(live bool) string {
	if live {
		return "live"
	}

	return "dummy"
}
