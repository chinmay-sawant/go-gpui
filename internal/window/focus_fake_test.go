package window

import (
	"context"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

// focusScreen is a fakeScreen with a focus order.
type focusScreen struct {
	*fakeScreen
	ids   []string
	at    int
	focus string
	nexts int
	prevs int
}

func (f *focusScreen) FocusNext(context.Context) error {
	f.nexts++
	if len(f.ids) == 0 {
		return nil
	}

	if f.focus == "" {
		f.at = 0
	} else {
		f.at = (f.at + 1) % len(f.ids)
	}

	f.focus = f.ids[f.at]

	return nil
}

func (f *focusScreen) FocusPrev(context.Context) error {
	f.prevs++
	if len(f.ids) == 0 {
		return nil
	}

	if f.focus == "" {
		f.at = len(f.ids) - 1
	} else {
		f.at = (f.at + len(f.ids) - 1) % len(f.ids)
	}

	f.focus = f.ids[f.at]

	return nil
}

func (f *focusScreen) Focus(_ context.Context, id string) error {
	f.focus = id

	return nil
}

func (f *focusScreen) FocusID() string { return f.focus }

var _ host.Focuser = (*focusScreen)(nil)
