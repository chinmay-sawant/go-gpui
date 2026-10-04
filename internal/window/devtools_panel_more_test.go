package window

import (
	"context"
	"strings"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

// TestDevToolsToggleSwallowsCtrlShiftI covers the second toggle chord.
func TestDevToolsToggleSwallowsCtrlShiftI(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	mods := modifiers{Control: true, Shift: true}

	swallow, err := s.devKeyStep(ebiten.KeyI, true, mods)
	if err != nil || !swallow {
		t.Fatalf("ctrl+shift+i swallow = %v, err = %v", swallow, err)
	}

	if d.on {
		t.Fatal("ctrl+shift+i did not close the overlay")
	}

	if down, up := s.pageKeyStep(ebiten.KeyI, true); down || up {
		t.Fatal("the page saw the ctrl+shift+i press")
	}
}

// TestDevHoverNeverCallsThePage pins the hover freeze while the overlay is
// on: the shell records the box, the page never sees a hover.
func TestDevHoverNeverCallsThePage(t *testing.T) {
	t.Parallel()

	d := newDevScreen()
	s := newDevShell(d)
	s.dev.on = true
	s.devHover(5, 5)

	if d.hovers != 0 {
		t.Fatalf("page hovers = %d, want 0 while the overlay is on", d.hovers)
	}
}

// TestDevPanelMovesWithARedraw checks the Phase 5 exit: the panel reads the
// page stats, so a redraw moves the redraw line in the same frame.
func TestDevPanelMovesWithARedraw(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p, err := page.New(page.Config{HTML: `<p id="a">one</p>`, Width: 200, Height: 100})
	if err != nil {
		t.Fatal(err)
	}

	if err = p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	s := &shell{app: p, ctx: ctx}
	s.dev.stats = p.Stats()

	if !strings.Contains(strings.Join(s.devLines(), "\n"), "redraws 1") {
		t.Fatalf("panel lines = %q", strings.Join(s.devLines(), "\n"))
	}

	if err = p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	s.dev.stats = p.Stats()

	if !strings.Contains(strings.Join(s.devLines(), "\n"), "redraws 2") {
		t.Fatalf("panel lines = %q", strings.Join(s.devLines(), "\n"))
	}
}
