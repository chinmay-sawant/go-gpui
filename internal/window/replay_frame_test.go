package window

import "testing"

// TestFrameDirtyBlits checks a ticking page that opted into dirty rects
// builds the buffer once and then blits it.
func TestFrameDirtyBlits(t *testing.T) {
	t.Parallel()

	screen := &dirtyScreen{fakeScreen: &fakeScreen{width: 100, height: 100}}
	screen.display = partialDisplay(100, 100)
	screen.gen = 1
	screen.ticking = true
	screen.frame = true
	s := newDirtyShell(screen)
	dst := partialBuf(100, 100)

	s.drawReplayPartial(dst)

	if s.partial.buf == nil {
		t.Fatal("opted-in tick did not build a buffer")
	}

	s.drawReplayPartial(dst)

	if s.partial.mode != repaintBlit {
		t.Fatalf("second frame mode = %v, want blit", s.partial.mode)
	}
}
