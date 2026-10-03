package window

import (
	"context"
	"errors"
	"testing"
)

// tickScreen is a fakeScreen with a per-frame callback.
type tickScreen struct {
	*fakeScreen
	ticks int
	err   error
}

func (s *tickScreen) Tick(context.Context) error {
	s.ticks++

	return s.err
}

func TestTickFrameRunsTicker(t *testing.T) {
	t.Parallel()

	screen := &tickScreen{fakeScreen: &fakeScreen{}}
	s := &shell{app: screen, ctx: context.Background()}

	if err := s.tickFrame(); err != nil {
		t.Fatalf("tickFrame: %v", err)
	}

	if screen.ticks != 1 {
		t.Fatalf("ticks = %d, want 1", screen.ticks)
	}
}

func TestTickFrameSkipsPlainScreen(t *testing.T) {
	t.Parallel()

	s := &shell{app: &fakeScreen{}, ctx: context.Background()}
	if err := s.tickFrame(); err != nil {
		t.Fatalf("tickFrame: %v", err)
	}
}

func TestTickFrameReturnsError(t *testing.T) {
	t.Parallel()

	want := errors.New("frame failed")
	screen := &tickScreen{fakeScreen: &fakeScreen{}, err: want}
	s := &shell{app: screen, ctx: context.Background()}

	if err := s.tickFrame(); !errors.Is(err, want) {
		t.Fatalf("err = %v", err)
	}
}
