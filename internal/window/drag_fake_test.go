package window

import (
	"context"
)

// selectScreen records the range calls the window makes.
type selectScreen struct {
	*fakeScreen
	calls []string
}

func (f *selectScreen) Press(context.Context, float64, float64) error {
	f.calls = append(f.calls, "press")

	return nil
}

func (f *selectScreen) Click(context.Context, float64, float64) error {
	f.calls = append(f.calls, "click")

	return nil
}

func (f *selectScreen) Release(context.Context) error {
	f.calls = append(f.calls, "release")

	return nil
}

func (f *selectScreen) SelectAt(context.Context, float64, float64) error {
	f.calls = append(f.calls, "at")

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
