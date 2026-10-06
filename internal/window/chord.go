package window

import "strings"

import "github.com/chinmay-sawant/ownframe/internal/clipboard"

func (s *shell) applyChord(chord chord) error {
	switch chord {
	case chordCopy:
		return s.copy(false)
	case chordCut:
		return s.copy(true)
	case chordPaste:
		text := strings.ReplaceAll(clipboard.Read(), "\r", "")
		text = strings.ReplaceAll(text, "\n", "")

		return s.app.Paste(s.ctx, text)
	case chordSelectAll:
		return s.app.SelectAll(s.ctx)
	case chordUndo:
		return s.app.Undo(s.ctx)
	case chordRedo:
		return s.app.Redo(s.ctx)
	default:
		return nil
	}
}

func (s *shell) copy(cut bool) error {
	var text string
	var ok bool
	var err error

	if cut {
		text, ok, err = s.app.Cut(s.ctx)
	} else {
		text, ok, err = s.app.Copy(s.ctx)
	}

	if err != nil || !ok || text == "" {
		return err
	}

	clipboard.Write(text)

	return nil
}
