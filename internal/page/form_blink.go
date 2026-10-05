package page

import (
	"context"
	"time"
)

// caretBlinkInterval is how long the caret stays shown or hidden.
const caretBlinkInterval = 530 * time.Millisecond

// resetCaretBlink shows the caret and restarts the blink clock. Every caret
// move and edit calls it, so the caret is solid while the person types.
func (p *Page) resetCaretBlink() {
	if p.form == nil {
		return
	}

	p.form.blinkOn = true
	p.form.blinkAt = time.Time{}
}

// caretBlink toggles the focused field's caret on the blink clock and draws
// when the phase changes. A range selection paints no caret, so it does
// nothing.
func (p *Page) caretBlink(ctx context.Context) error {
	c, ok := p.focusedEditable()
	if !ok {
		return nil
	}

	n := runeLen(c.Value)
	start, end := p.form.bounds(n)
	if start != end {
		return nil
	}

	now := time.Now()
	if p.form.blinkAt.IsZero() {
		p.form.blinkAt = now
		p.form.blinkOn = true

		return nil
	}

	if now.Sub(p.form.blinkAt) < caretBlinkInterval {
		return nil
	}

	p.form.blinkOn = !p.form.blinkOn
	p.form.blinkAt = now

	return p.Redraw(ctx)
}
