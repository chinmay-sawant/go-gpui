package window

import (
	"context"
	"time"
)

// selector is the page range surface the window drives while the pointer is
// down. A page without it keeps the old click behavior.
type selector interface {
	SelectAt(ctx context.Context, x, y float64) error
	Drag(ctx context.Context, x, y float64) error
	SelectWordAt(ctx context.Context, x, y float64) error
	SelectLineAt(ctx context.Context, x, y float64) error
}

// pressAt sends a press, places the caret, then sends the click. The caret
// lands before the click handler runs, so focus and a selection that the
// handler sets are not blurred afterwards. A second press selects the word
// under the point and a third selects the line.
func (s *shell) pressAt(px, py float64) error {
	if err := s.app.Press(s.ctx, px, py); err != nil {
		return err
	}

	sel, ok := s.app.(selector)
	if !ok {
		return s.app.Click(s.ctx, px, py)
	}

	if err := sel.SelectAt(s.ctx, px, py); err != nil {
		return err
	}

	if err := s.app.Click(s.ctx, px, py); err != nil {
		return err
	}

	s.dragActive = true
	s.dragX, s.dragY = px, py

	count := s.clicks.step(time.Now(), px, py)
	if count == 2 {
		return sel.SelectWordAt(s.ctx, px, py)
	}

	if count >= 3 {
		return sel.SelectLineAt(s.ctx, px, py)
	}

	return nil
}

// dragAt extends the range while the button stays down and the point moves.
func (s *shell) dragAt(px, py float64) error {
	if !s.dragActive || (px == s.dragX && py == s.dragY) {
		return nil
	}

	sel, ok := s.app.(selector)
	if !ok {
		return nil
	}

	s.dragX, s.dragY = px, py

	return sel.Drag(s.ctx, px, py)
}

// releaseAt ends a press and stops the drag.
func (s *shell) releaseAt() error {
	s.dragActive = false

	return s.app.Release(s.ctx)
}
