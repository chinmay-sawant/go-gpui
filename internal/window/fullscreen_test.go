package window

import (
	"testing"
)

func TestF11TogglesFullscreen(t *testing.T) {
	t.Parallel()

	full := false
	var applied []bool

	s := &shell{}
	s.readFullscreen = func() bool { return full }
	s.applyFullscreen = func(on bool) { applied = append(applied, on) }

	if !s.f11() {
		t.Fatal("f11 was not consumed")
	}

	if len(applied) != 1 || !applied[0] {
		t.Fatalf("applied = %v", applied)
	}

	full = true
	if !s.f11() {
		t.Fatal("second f11 was not consumed")
	}

	if len(applied) != 2 || applied[1] {
		t.Fatalf("applied = %v", applied)
	}
}
