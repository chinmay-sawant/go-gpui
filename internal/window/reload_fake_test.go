package window

import (
	"context"

	"github.com/chinmay-sawant/blinkless/layout"
)

// reloadScreen is a fake screen whose page reports one reload.
type reloadScreen struct {
	*fakeScreen
	polls   int
	left    int
	err     error
	display *layout.Display
	hovers  int
	hoverX  float64
	hoverY  float64
}

func (r *reloadScreen) Hover(_ context.Context, x, y float64) error {
	r.hovers++
	r.hoverX, r.hoverY = x, y

	return nil
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
