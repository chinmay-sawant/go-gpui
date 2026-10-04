package window

import "testing"

func TestPartialPathNeedsATaker(t *testing.T) {
	t.Parallel()

	plain := &fakeScreen{width: 100, height: 100}
	s := &shell{
		app: plain, display: partialDisplay(100, 100), seq: 1,
		screenW: 100, screenH: 100,
	}

	s.drawReplayPartial(partialBuf(100, 100))

	if s.partial.buf != nil {
		t.Fatal("buffer built without a dirty taker")
	}
}

func TestTickingPageKeepsFullReplay(t *testing.T) {
	t.Parallel()

	screen := &dirtyScreen{fakeScreen: &fakeScreen{width: 100, height: 100}}
	screen.display = partialDisplay(100, 100)
	screen.ticking = true
	s := newDirtyShell(screen)

	s.drawReplayPartial(partialBuf(100, 100))

	if screen.takes != 0 {
		t.Fatalf("TakeDirty calls = %d, want 0", screen.takes)
	}

	if s.partial.buf != nil {
		t.Fatal("buffer built for a ticking page")
	}
}

func TestScrollBlitsTheBuffer(t *testing.T) {
	t.Parallel()

	screen := &dirtyScreen{fakeScreen: &fakeScreen{width: 100, height: 100}}
	screen.display = partialDisplay(100, 100)
	screen.gen = 1
	s := newDirtyShell(screen)
	dst := partialBuf(100, 100)

	s.drawReplayPartial(dst)
	s.scrollY = 20
	screen.ok = false
	s.drawReplayPartial(dst)

	if s.partial.mode != repaintBlit {
		t.Fatalf("scroll mode = %v, want blit", s.partial.mode)
	}
}
