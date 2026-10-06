package player

import (
	"time"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/internal/frame"
)

// accent is Spotify green, #1db954.
var accent = [3]float64{29.0 / 255, 185.0 / 255, 84.0 / 255}

// animate moves the seek bar, the clock, the credit, and the equalizer
// without another Redraw.
func (a *App) animate() {
	d := a.page.Display()
	if d == nil {
		return
	}

	boxes := a.page.Boxes()
	pos := a.audio.Position()
	dur := a.audio.Duration()

	if dur > 0 {
		a.animateSeek(d, boxes, pos, dur)
		a.animateTimes(d, boxes, pos, dur)
		a.updateView(pos, dur)
	}

	a.animateCredit(d, boxes)
	a.animateEq(d, boxes)
}

// animateSeek stretches the accent fill to the playing fraction.
func (a *App) animateSeek(d *ownframe.Display, boxes []ownframe.Box, pos, dur time.Duration) {
	seek, ok := boxByID(boxes, "seek")
	if !ok {
		return
	}

	fill := frame.Fill(d, seek, accent)
	if fill == nil {
		return
	}

	_, _, w, _ := frame.BoxUnits(d, seek)
	fraction := float64(pos) / float64(dur)
	if fraction < 0 {
		fraction = 0
	}

	if fraction > 1 {
		fraction = 1
	}

	fill.W = w * fraction
}

// animateTimes writes the elapsed and remaining labels when they change.
func (a *App) animateTimes(d *ownframe.Display, boxes []ownframe.Box, pos, dur time.Duration) {
	if box, ok := boxByID(boxes, "elapsed"); ok {
		if op := frame.Text(d, box); op != nil {
			if text := formatSeconds(int(pos.Seconds())); op.Text != text {
				op.Text = text
			}
		}
	}

	if box, ok := boxByID(boxes, "remaining"); ok {
		if op := frame.Text(d, box); op != nil {
			rem := int((dur - pos).Seconds())
			if rem < 0 {
				rem = 0
			}

			if text := "-" + formatSeconds(rem); op.Text != text {
				op.Text = text
			}
		}
	}
}

// boxByID returns the first hit-test box with id.
func boxByID(boxes []ownframe.Box, id string) (ownframe.Box, bool) {
	for _, box := range boxes {
		if box.ID == id {
			return box, true
		}
	}

	return ownframe.Box{}, false
}
