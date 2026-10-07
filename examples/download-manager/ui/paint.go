package ui

import (
	"fmt"
	"strconv"

	"github.com/chinmay-sawant/ownframe"
)

// paint edits the retained operations for the current view. It does nothing
// on the bitmap fallback path, where there is no display list.
func (a *App) paint() {
	if a.page.Display() == nil {
		return
	}

	if a.page.Generation() != a.bound {
		a.bind()
	}

	for i, row := range a.view.Active {
		if i >= len(a.bindings.fill) {
			break
		}

		if op := a.bindings.fill[i]; op != nil {
			f := row.Fraction()
			if f <= 0 {
				op.W = 0
			} else {
				op.W = a.bindings.trackW[i] * f
			}
		}

		setText(a.bindings.pct[i], row.Progress())
		setText(a.bindings.speed[i], row.SpeedText())
		setText(a.bindings.eta[i], row.ETAText())
	}

	setText(a.bindings.sum[0], strconv.Itoa(a.view.Summary.Active()))
	setText(a.bindings.sum[1], strconv.Itoa(a.view.Summary.Running))
	setText(a.bindings.sum[2], strconv.Itoa(a.view.Summary.Completed))
	setText(a.bindings.sum[3], strconv.Itoa(a.view.Summary.Failed))

	if detail := a.view.Detail; detail != nil {
		setText(a.bindings.detail[0], detail.StateText())
		setText(a.bindings.detail[1], detail.SizeText())
		setText(a.bindings.detail[2], detail.SpeedText())
	}

	setText(a.bindings.foot, a.footText())
}

// setText writes a retained text run only when the value changed.
func setText(op *ownframe.DisplayOp, text string) {
	if op != nil && op.Text != text {
		op.Text = text
	}
}

// footText is the measured tick and repaint line.
func (a *App) footText() string {
	s := a.view.Stats

	return fmt.Sprintf("tick %s · paints %d · redraws %d · applied %d",
		s.TickText(), s.Paints, s.Redraws, s.Applied)
}
