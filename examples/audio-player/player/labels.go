package player

import (
	"time"

	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/internal/frame"
)

// timeText rewrites the elapsed and remaining labels.
func (a *App) timeText(d *gpui.Display, pos, dur time.Duration) {
	if dur <= 0 {
		return
	}

	setText(d, a.page.Boxes(), "elapsed", clock(int(pos.Seconds())))
	setText(d, a.page.Boxes(), "remaining", "-"+clock(int((dur-pos).Seconds())))
}

// setText changes one labelled text op when the string differs.
func setText(d *gpui.Display, boxes []gpui.Box, id, text string) {
	box, ok := boxByID(boxes, id)
	if !ok {
		return
	}

	if op := frame.Text(d, box); op != nil && op.Text != text {
		op.Text = text
	}
}

// creditText shows the loading note or the clip credit and keeps it in the
// view, so a later Redraw prints the same line.
func (a *App) creditText(d *gpui.Display) {
	text := a.view.Credit

	switch {
	case a.audio.Loading():
		text = "Loading free audio..."
	case a.audio.Credit().Title != "":
		c := a.audio.Credit()
		text = c.Title + " - " + c.Creator + " (" + c.License + ")"
	case a.audio.Err() != nil:
		text = "Audio unavailable"
	}

	a.view.Credit = text

	box, ok := boxByID(a.page.Boxes(), "credit")
	if !ok {
		return
	}

	if op := frame.Text(d, box); op != nil && op.Text != text {
		op.Text = text
	}
}
