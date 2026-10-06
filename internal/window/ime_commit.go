//go:build android || ios

package window

import (
	"github.com/hajimehoshi/ebiten/v2/exp/textinput"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

// imeCommit applies a committed edit.
func (s *shell) imeCommit(c *textinput.Commit) {
	s.imeDrop()

	if before, after := c.IsSurroundingTextReplaced(); before || after {
		if target, ok := s.app.(host.IME); ok {
			prefix, suffix := c.SurroundingText()
			text := c.Text()
			caret := len(prefix) + len(text)
			err := target.IMEReplace(s.ctx, s.ime.start, s.ime.end, prefix+text+suffix, caret)
			if err != nil {
				s.ime.err = err
			}
		}

		return
	}

	if c.Text() == "" {
		return
	}

	if err := s.app.Type(s.ctx, c.Text()); err != nil {
		s.ime.err = err
	}
}

// imeEndByUser hides the keyboard by dropping the focus.
func (s *shell) imeEndByUser() {
	s.ime.last = ""

	if f, ok := s.app.(host.Focuser); ok {
		if err := f.Focus(s.ctx, ""); err != nil {
			s.ime.err = err
		}
	}
}
