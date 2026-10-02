package window

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestPressedNow(t *testing.T) {
	t.Parallel()

	cases := []struct {
		down bool
		was  bool
		want bool
	}{
		{false, false, false},
		{true, false, true},
		{true, true, false},
		{false, true, false},
	}

	for _, c := range cases {
		if got := pressedNow(c.down, c.was); got != c.want {
			t.Fatalf("pressedNow(%v, %v) = %v", c.down, c.was, got)
		}
	}
}

func TestFreshTouches(t *testing.T) {
	t.Parallel()

	got := freshTouches(
		[]ebiten.TouchID{1, 2, 5},
		[]ebiten.TouchID{2, 3},
	)
	if len(got) != 2 || got[0] != 1 || got[1] != 5 {
		t.Fatalf("fresh = %v", got)
	}

	if freshTouches(nil, []ebiten.TouchID{1}) != nil {
		t.Fatal("fresh from nil is not nil")
	}
}
