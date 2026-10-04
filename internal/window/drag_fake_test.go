package window

import (
	"context"
)

// selectScreen records the range calls the window makes. focus models the
// state a click handler sets and SelectAt clears.
type selectScreen struct {
	*fakeScreen
	calls []string
	focus bool
}

func (f *selectScreen) Press(context.Context, float64, float64) error {
	f.calls = append(f.calls, "press")

	return nil
}

func (f *selectScreen) Click(context.Context, float64, float64) error {
	f.calls = append(f.calls, "click")
	f.focus = true

	return nil
}

func (f *selectScreen) Release(context.Context) error {
	f.calls = append(f.calls, "release")

	return nil
}

func (f *selectScreen) SelectAt(context.Context, float64, float64) error {
	f.calls = append(f.calls, "at")
	f.focus = false

	return nil
}

func (f *selectScreen) Drag(context.Context, float64, float64) error {
	f.calls = append(f.calls, "drag")

	return nil
}

func (f *selectScreen) SelectWordAt(context.Context, float64, float64) error {
	f.calls = append(f.calls, "word")

	return nil
}

func (f *selectScreen) SelectLineAt(context.Context, float64, float64) error {
	f.calls = append(f.calls, "line")

	return nil
}
