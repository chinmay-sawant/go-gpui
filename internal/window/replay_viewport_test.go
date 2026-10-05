package window

import (
	"image"
	"testing"
)

func TestViewportBufferStaysWindowSized(t *testing.T) {
	app := &dirtyScreen{fakeScreen: &fakeScreen{width: 80, height: 60}, gen: 1}
	app.display = partialDisplay(80, 37000)
	s := newDirtyShell(app)
	dst := partialBuf(80, 60)
	s.partial.buf = partialBuf(80, 100)
	s.drawViewport(dst, app.display, app)
	if got := s.viewport.buf.Bounds(); got != image.Rect(0, 0, 80, 180) {
		t.Fatalf("buffer = %v", got)
	}
	if s.partial.buf != nil {
		t.Fatal("content buffer retained beside viewport cache")
	}
	saved := s.viewport.buf
	s.drawViewport(dst, app.display, app)
	if s.viewport.buf != saved {
		t.Fatal("idle frame allocated another viewport")
	}
	s.scrollY = 2000
	s.drawViewport(dst, app.display, app)
	if s.viewport.area.Min.Y > 2000 || s.viewport.area.Max.Y < 2060 || s.viewport.buf != saved {
		t.Fatal("scroll replaced the buffer or kept stale pixels")
	}
	app.gen++
	s.seq = app.gen
	s.drawViewport(dst, app.display, app)
	if s.viewport.key.generation != app.gen {
		t.Fatal("generation did not invalidate viewport")
	}
	s.disposeViewport()
	if s.viewport.buf != nil || s.viewport.valid {
		t.Fatal("viewport not released")
	}
}
