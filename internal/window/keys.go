package window

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

func (s *shell) keys() error {
	mods := readModifiers()

	if err := s.devtoolsKeys(mods); err != nil {
		return err
	}

	if err := s.keyEvents(); err != nil {
		return err
	}

	if typingSuppressed(mods) && backspaceDue() {
		return s.app.DeleteWord(s.ctx)
	}

	if chord := s.chords.pressed(mods); chord != chordNone {
		return s.applyChord(chord)
	}

	if typingSuppressed(mods) {
		return s.submitIfEnter()
	}

	s.chars = ebiten.AppendInputChars(s.chars[:0])
	s.chars = s.chords.filter(s.chars)
	text := strings.ReplaceAll(string(s.chars), "\r", "")
	text = strings.ReplaceAll(text, "\n", "")

	if text != "" {
		if err := s.app.Type(s.ctx, text); err != nil {
			return err
		}
	}

	if backspaceDue() {
		if err := s.app.Backspace(s.ctx); err != nil {
			return err
		}
	}

	return s.submitIfEnter()
}

func readModifiers() modifiers {
	return modifiers{
		Control: ebiten.IsKeyPressed(ebiten.KeyControl),
		Meta:    ebiten.IsKeyPressed(ebiten.KeyMeta),
		Alt:     ebiten.IsKeyPressed(ebiten.KeyAlt),
		Shift:   ebiten.IsKeyPressed(ebiten.KeyShift),
	}
}
