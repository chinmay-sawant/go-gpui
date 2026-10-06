// Package scene draws the tetris example on ownframe pages: a game screen
// with a 10x20 board of retained fills, a next-piece preview, the score
// panel, and the phase overlays, plus a paged score-history screen. The
// frame callback steps the core game through game.Clock and repaints the
// retained operations, so routine frames never parse the HTML again. A
// Redraw happens on a resize, a theme change, a phase overlay, and when
// the history screen opens or closes.
package scene

import (
	"context"
	"time"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// Scene is the tetris screen and its frame loop.
type Scene struct {
	page  *ownframe.Page
	model Model
	store Store

	now     func() time.Time
	step    Stepper
	focus   func() bool
	focused bool
	last    time.Time

	ops   ops
	bound uint64

	frame    Frame
	dirty    bool
	overlays bool
	bitmap   time.Time
	runID    string

	dark   bool
	status string
	auto   bool

	settingsReq, themeReq, scoresReq, demoReq, saveReq, snapReq uint64
	pendingSave                                                 *game.Result
	pendingReplay                                               *game.Replay
	nextSaveTry                                                 time.Time

	hist history
}

// Page returns the page Run displays.
func (s *Scene) Page() *ownframe.Page { return s.page }

// Redraw renders the current frame through the template.
func (s *Scene) Redraw(ctx context.Context) error { return s.redraw(ctx) }

// redraw pushes the current view into the page.
func (s *Scene) redraw(ctx context.Context) error {
	s.page.SetData(s.view())
	s.dirty = false

	return s.page.Redraw(ctx)
}

// view packs the current frame and UI state for the template.
func (s *Scene) view() view { return buildView(s.frame, s.dark, s.status, &s.hist) }

// setStatus changes the status line and marks the picture stale.
func (s *Scene) setStatus(text string) {
	s.status = text
	s.dirty = true
}
