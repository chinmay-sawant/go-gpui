package window

import (
	"context"
	"image"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// dirtyScreen is a fakeScreen that keeps a display list and reports one
// scripted dirty rect, the way a page with TakeDirty does.
type dirtyScreen struct {
	*fakeScreen
	display *layout.Display
	gen     uint64
	layouts int
	rect    image.Rectangle
	ok      bool
	takes   int
	ticking bool
}

func (d *dirtyScreen) Display() *layout.Display { return d.display }
func (d *dirtyScreen) Generation() uint64       { return d.gen }
func (d *dirtyScreen) Ticking() bool            { return d.ticking }

func (d *dirtyScreen) TakeDirty() (image.Rectangle, bool) {
	d.takes++

	return d.rect, d.ok
}

// click stands in for one counter click: the page lays out once and reports
// one changed box.
func (d *dirtyScreen) click(rect image.Rectangle) {
	d.gen++
	d.layouts++
	d.rect = rect
	d.ok = true
}

func newDirtyShell(screen *dirtyScreen) *shell {
	return &shell{
		app: screen, ctx: context.Background(),
		display: screen.display, seq: screen.gen,
		screenW: screen.width, screenH: screen.height,
	}
}
