package window

import "testing"

// TestDevClampScroll checks the content scroll stays within the rows.
func TestDevClampScroll(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.dev.rows = make([]devRow, 200)
	s.dev.content = devRect{H: 100}

	s.dev.scroll = -5
	s.devClampScroll()

	if s.dev.scroll != 0 {
		t.Fatalf("negative scroll = %v, want 0", s.dev.scroll)
	}

	s.dev.scroll = 1e9
	s.devClampScroll()

	want := float64(200)*devLineHeight() - 100
	if s.dev.scroll != want {
		t.Fatalf("clamped scroll = %v, want %v", s.dev.scroll, want)
	}
}

// TestDevClampScrollShortContent checks a tab shorter than the viewport
// cannot scroll.
func TestDevClampScrollShortContent(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.dev.rows = make([]devRow, 2)
	s.dev.content = devRect{H: 400}
	s.dev.scroll = 30

	s.devClampScroll()

	if s.dev.scroll != 0 {
		t.Fatalf("short content scroll = %v, want 0", s.dev.scroll)
	}
}
