package window

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/host"
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// scrollScreen is a fakeScreen with one pending scroll request.
type scrollScreen struct {
	*fakeScreen
	req host.Scroll
	has bool
}

func (f *scrollScreen) TakeScroll() (host.Scroll, bool) {
	if !f.has {
		return host.Scroll{}, false
	}

	f.has = false

	return f.req, true
}

func TestScrollRequestClamps(t *testing.T) {
	t.Parallel()

	app := &scrollScreen{
		fakeScreen: &fakeScreen{
			width: 200, height: 200,
			boxes: []layout.Box{{X: 0, Y: 0, W: 400, H: 500}},
		},
		req: host.Scroll{Y: 1000, Absolute: true},
		has: true,
	}
	s := &shell{app: app, ctx: context.Background(), screenW: 200, screenH: 200}

	s.applyScrollRequest()

	if s.scrollX != 0 || s.scrollY != 300 {
		t.Fatalf("absolute = %d, %d", s.scrollX, s.scrollY)
	}

	app.req = host.Scroll{X: 10, Y: -40}
	app.has = true
	s.applyScrollRequest()

	if s.scrollX != 10 || s.scrollY != 260 {
		t.Fatalf("delta = %d, %d", s.scrollX, s.scrollY)
	}

	s.applyScrollRequest()

	if s.scrollY != 260 {
		t.Fatalf("cleared request moved to %d", s.scrollY)
	}
}

func TestScrollRequestPlainScreen(t *testing.T) {
	t.Parallel()

	s := &shell{app: &fakeScreen{}, ctx: context.Background()}
	s.applyScrollRequest()
}
