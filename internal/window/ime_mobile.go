//go:build android || ios

package window

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2/exp/textinput"

	"github.com/chinmay-sawant/ownframe/internal/host"
)

// imeState is the platform text input session for the focused control.
type imeState struct {
	composer textinput.Composer
	field    string
	last     string
	caret    image.Rectangle
	start    int
	end      int
	handled  bool
	err      error
}

// imeInit wires the composer callbacks.
func (s *shell) imeInit() {
	s.ime.composer.OnNewSession = s.imeNewSession
	s.ime.composer.OnComposition = s.imeComposition
	s.ime.composer.OnCommit = s.imeCommit
	s.ime.composer.OnEndByUser = s.imeEndByUser
}

// imeHandled reports whether the platform input consumed this tick.
func (s *shell) imeHandled() bool { return s.ime.handled }

// imeUpdate runs the text input session for one tick.
func (s *shell) imeUpdate() error {
	id := ""
	if f, ok := s.app.(host.Focuser); ok {
		id = f.FocusID()
	}

	if id != s.ime.field {
		if s.ime.field != "" {
			s.ime.composer.Cancel()
			s.ime.last = ""
		}

		s.ime.field = id
	}

	s.ime.err = nil
	s.imeResize()
	handled, err := s.ime.composer.Update()
	s.ime.handled = handled

	if err != nil {
		return err
	}

	return s.ime.err
}

// imeNewSession reports the caret to the platform when a text control is
// focused.
func (s *shell) imeNewSession() *textinput.SessionOptions {
	target, ok := s.app.(host.IME)
	if !ok {
		return nil
	}

	box, caret, before, after, ok := target.IMEContext()
	if !ok {
		return nil
	}

	s.ime.start = caret - len(before)
	s.ime.end = caret + len(after)
	s.ime.caret = imeCaretRect(box)

	return &textinput.SessionOptions{
		CaretBounds:     s.ime.caret,
		TextBeforeCaret: before,
		TextAfterCaret:  after,
	}
}
