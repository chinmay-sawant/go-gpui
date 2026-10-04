package window

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestKeyEventNamePrefixesCaretKeys(t *testing.T) {
	t.Parallel()

	cases := []struct {
		key  ebiten.Key
		mods modifiers
		want string
	}{
		{ebiten.KeyArrowLeft, modifiers{}, "arrowleft"},
		{ebiten.KeyArrowLeft, modifiers{Shift: true}, "shift+arrowleft"},
		{ebiten.KeyArrowRight, modifiers{Control: true}, "ctrl+arrowright"},
		{ebiten.KeyArrowRight, modifiers{Alt: true}, "alt+arrowright"},
		{ebiten.KeyHome, modifiers{Control: true, Shift: true}, "ctrl+shift+home"},
		{ebiten.KeyEnd, modifiers{Meta: true}, "ctrl+end"},
		{ebiten.KeyA, modifiers{Control: true}, "a"},
		{ebiten.KeySpace, modifiers{Shift: true}, "space"},
	}

	for _, c := range cases {
		if got := keyEventName(c.key, c.mods); got != c.want {
			t.Fatalf("keyEventName(%v, %+v) = %q, want %q", c.key, c.mods, got, c.want)
		}
	}
}
