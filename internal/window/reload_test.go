package window

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// reloadScreen is a fake screen whose page reports one reload.
type reloadScreen struct {
	*fakeScreen
	polls   int
	left    int
	err     error
	display *layout.Display
}

func newReloadScreen() *reloadScreen {
	return &reloadScreen{
		fakeScreen: &fakeScreen{width: 80, height: 60, minW: 1, minH: 1, redraws: 1},
		display:    &layout.Display{Width: 80, Height: 60},
	}
}

func (r *reloadScreen) PollReload(context.Context) (bool, error) {
	r.polls++

	if r.err != nil {
		return false, r.err
	}

	if r.left <= 0 {
		return false, nil
	}

	r.left--
	r.redraws++
	r.boxes = nil
	r.display = &layout.Display{Width: 80, Height: 60}

	return true, nil
}

func (r *reloadScreen) Display() *layout.Display { return r.display }

func TestHotReloadPollsOnTheInterval(t *testing.T) {
	t.Parallel()

	app := newReloadScreen()
	app.left = 1

	s := NewGame(context.Background(), app).(*shell)
	if err := s.Update(); err != nil {
		t.Fatal(err)
	}

	if app.polls != 1 {
		t.Fatalf("polls = %d, want 1", app.polls)
	}

	if app.redraws != 2 {
		t.Fatalf("redraws = %d, want one reload", app.redraws)
	}

	if s.display == nil || s.seq != 2 {
		t.Fatalf("display = %v, seq = %d, want the new generation", s.display, s.seq)
	}
}

func TestHotReloadWaitsForTheInterval(t *testing.T) {
	t.Parallel()

	app := newReloadScreen()
	app.left = 1

	s := NewGame(context.Background(), app).(*shell)
	s.lastPoll = time.Now()

	if err := s.Update(); err != nil {
		t.Fatal(err)
	}

	if app.polls != 0 {
		t.Fatalf("polls = %d, want 0 inside the interval", app.polls)
	}

	s.lastPoll = time.Now().Add(-time.Second)

	if err := s.Update(); err != nil {
		t.Fatal(err)
	}

	if app.polls != 1 {
		t.Fatalf("polls = %d, want 1 after the interval", app.polls)
	}
}
