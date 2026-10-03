package player

import (
	"time"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/internal/frame"
)

// accent is #6d28d9, the player's violet.
var accent = [3]float64{109.0 / 255, 40.0 / 255, 217.0 / 255}

// animate paints the moving parts from the audio position without a Redraw.
func (a *App) animate() {
	d := a.page.Display()
	if d == nil {
		return
	}

	pos := a.audio.Position()
	dur := a.audio.Duration()

	a.seekFill(d, pos, dur)
	a.timeText(d, pos, dur)
	a.creditText(d)
	a.equalizer(d, a.audio.Playing())
	a.syncView(pos, dur)
}

// seekFill stretches the accent fill across the seek bar.
func (a *App) seekFill(d *gpui.Display, pos, dur time.Duration) {
	if dur <= 0 {
		return
	}

	box, ok := boxByID(a.page.Boxes(), "seek")
	if !ok {
		return
	}

	fill := frame.Fill(d, box, accent)
	if fill == nil {
		return
	}

	_, _, w, _ := frame.BoxUnits(d, box)
	fill.W = w * float64(pos) / float64(dur)
}

// syncView keeps the view's clock in step with the audio, so the next
// click-driven Redraw prints current values.
func (a *App) syncView(pos, dur time.Duration) {
	if dur <= 0 {
		return
	}

	a.view.Progress = clamp(int(float64(pos)/float64(dur)*100), 0, 100)
	a.view.Elapsed = clock(int(pos.Seconds()))
	a.view.Remaining = "-" + clock(int((dur - pos).Seconds()))
}

// boxByID finds one hit-test box by id.
func boxByID(boxes []gpui.Box, id string) (gpui.Box, bool) {
	for _, box := range boxes {
		if box.ID == id {
			return box, true
		}
	}

	return gpui.Box{}, false
}
