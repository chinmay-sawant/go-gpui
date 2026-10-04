package window

import (
	"context"
	"testing"
	"time"
)

func TestResizeRedrawsOnTheFirstChangedFrame(t *testing.T) {
	t.Parallel()

	app := &fakeScreen{
		width: 480, height: 640,
		minW: 320, minH: 400,
		maxW: 2560, maxH: 2560,
	}
	game := NewGame(context.Background(), app)
	shell, ok := game.(*shell)
	if !ok {
		t.Fatalf("game = %T", game)
	}

	shell.Layout(600, 720)
	if err := shell.resize(); err != nil {
		t.Fatal(err)
	}

	if app.redraws != 1 {
		t.Fatalf("redraws = %d, want 1", app.redraws)
	}

	if app.width != 600 || app.height != 720 {
		t.Fatalf("size = %d x %d", app.width, app.height)
	}

	if err := shell.resize(); err != nil {
		t.Fatal(err)
	}

	if app.redraws != 1 {
		t.Fatalf("redraws after a settled frame = %d, want 1", app.redraws)
	}
}

func TestResizeClampsBeforeRedraw(t *testing.T) {
	t.Parallel()

	app := &fakeScreen{
		width: 480, height: 640,
		minW: 320, minH: 400,
		maxW: 800, maxH: 900,
	}
	game := NewGame(context.Background(), app)
	shell := game.(*shell)

	shell.Layout(50, 50)
	if err := shell.resize(); err != nil {
		t.Fatal(err)
	}

	if app.width != 320 || app.height != 400 {
		t.Fatalf("clamped size = %d x %d", app.width, app.height)
	}

	if app.redraws != 1 {
		t.Fatalf("redraws = %d, want 1", app.redraws)
	}
}

func TestResizeCommitsOncePerSettledSize(t *testing.T) {
	t.Parallel()

	app := &fakeScreen{
		width: 480, height: 640,
		minW: 1, minH: 1,
		maxW: 2560, maxH: 2560,
	}
	game := NewGame(context.Background(), app)
	shell := game.(*shell)

	// Thirty window events between two updates settle on the last size.
	for i := 0; i < 30; i++ {
		shell.Layout(500+i, 700+i)
	}

	if err := shell.resize(); err != nil {
		t.Fatal(err)
	}

	if app.redraws != 1 {
		t.Fatalf("redraws = %d, want 1", app.redraws)
	}

	if app.width != 529 || app.height != 729 {
		t.Fatalf("size = %d x %d, want 529 x 729", app.width, app.height)
	}

	// Motion inside the throttle keeps the previous frame on screen.
	shell.lastRelayout = time.Now()

	for i := 0; i < 30; i++ {
		shell.Layout(600+i, 800+i)
	}

	if err := shell.resize(); err != nil {
		t.Fatal(err)
	}

	if app.redraws != 1 {
		t.Fatalf("throttled redraws = %d, want 1", app.redraws)
	}

	// The next update has no event, so the size has settled.
	if err := shell.resize(); err != nil {
		t.Fatal(err)
	}

	if app.redraws != 2 {
		t.Fatalf("settled redraws = %d, want 2", app.redraws)
	}

	if app.width != 629 || app.height != 829 {
		t.Fatalf("settled size = %d x %d, want 629 x 829", app.width, app.height)
	}
}
