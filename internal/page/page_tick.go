package page

import "context"

// SetTick registers fn to run once per frame before the window draws. fn can
// read the display list and change an operation's geometry or text, or call
// Redraw; either changes the next frame. nil removes the callback.
// Serve does not tick.
func (p *Page) SetTick(fn func(ctx context.Context) error) {
	p.tick = fn
}

// Tick runs the frame callback registered with SetTick and the caret blink.
// With no callback and no focused field it does nothing.
func (p *Page) Tick(ctx context.Context) error {
	if p.tick != nil {
		if err := p.tick(ctx); err != nil {
			return err
		}
	}

	return p.caretBlink(ctx)
}
