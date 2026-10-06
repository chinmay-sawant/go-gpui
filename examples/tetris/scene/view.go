package scene

import "github.com/chinmay-sawant/ownframe/examples/tetris/game"

// view is the template data for both screens; each screen uses the fields
// it needs. It is built only when the page is redrawn.
type view struct {
	Cells []cellV
	Next  []cellV
	Rows  []rowV

	Score  string
	Lines  string
	Level  string
	Status string
	Theme  string
	Title  string
	Page   string
	Msg    string
	Demo   string

	Ready, Paused, Over bool
	OverScore           string
}

// cellV places one template cell.
type cellV struct {
	ID, Class string
	Left, Top int
}

// rowV places one history row.
type rowV struct {
	Text string
	Top  int
}

// buildView packs the frame and the UI state for the template.
func buildView(f Frame, dark bool, status string, h *history) view {
	v := view{
		Cells:  boardCells(f),
		Next:   previewCells(f.Next),
		Rows:   h.rows(),
		Score:  padScore(f.Score),
		Lines:  padLines(f.Lines),
		Level:  padLevel(f.Level),
		Status: status,
		Theme:  themeLabel(dark),
		Title:  h.title(),
		Page:   pageLabel(h.page, h.more),
		Msg:    h.msg,
		Demo:   h.demoToggle(),
	}

	switch f.Phase {
	case game.PhaseReady:
		v.Ready = true
	case game.PhasePaused:
		v.Paused = true
	case game.PhaseOver:
		v.Over = true
		v.OverScore = "SCORE " + padScore(f.Score)
	}

	return v
}
