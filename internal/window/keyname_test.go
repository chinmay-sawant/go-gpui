package window

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestKeyName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		key  ebiten.Key
		want string
	}{
		{key: ebiten.KeySpace, want: "space"},
		{key: ebiten.KeyArrowUp, want: "arrowup"},
		{key: ebiten.KeyArrowDown, want: "arrowdown"},
		{key: ebiten.KeyEnter, want: "enter"},
		{key: ebiten.KeyEscape, want: "escape"},
		{key: ebiten.KeyA, want: "a"},
		{key: ebiten.KeyZ, want: "z"},
		{key: ebiten.KeyDigit1, want: "1"},
		{key: ebiten.Key0, want: "0"},
		{key: ebiten.KeyF5, want: "f5"},
		{key: ebiten.Key(9999), want: ""},
	}

	for _, test := range tests {
		if got := keyName(test.key); got != test.want {
			t.Errorf("keyName(%v) = %q, want %q", test.key, got, test.want)
		}
	}
}
