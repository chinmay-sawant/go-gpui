package window

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func pressedChord(mods modifiers) chord {
	keys := []ebiten.Key{
		ebiten.KeyC,
		ebiten.KeyV,
		ebiten.KeyX,
		ebiten.KeyA,
		ebiten.KeyZ,
		ebiten.KeyY,
		ebiten.KeyInsert,
		ebiten.KeyDelete,
	}

	for _, key := range keys {
		if !inpututil.IsKeyJustPressed(key) {
			continue
		}

		if chord := shortcutChord(mods, key); chord != chordNone {
			return chord
		}
	}

	return chordNone
}

func (s *shell) submitIfEnter() error {
	enter := inpututil.IsKeyJustPressed(ebiten.KeyEnter)
	numpad := inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter)
	if enter || numpad {
		return s.app.Submit(s.ctx)
	}

	return nil
}

func backspaceDue() bool {
	held := inpututil.KeyPressDuration(ebiten.KeyBackspace)
	if held == 1 {
		return true
	}

	if held <= backspaceRepeatAt {
		return false
	}

	return (held-backspaceRepeatAt)%backspaceRepeatEvery == 0
}
