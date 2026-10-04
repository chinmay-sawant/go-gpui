// Package input is the keyboard focus and pointer interaction example.
// Two fields, a checkbox, a link, and a long list put focus traversal, the
// caret keys, drag selection, the context menu, cursor shapes, touch scroll
// and pinch, page scrolling, and F11 in one window. gpui opens the window;
// this package does not.
package input

import (
	"context"
	_ "embed"

	"github.com/chinmay-sawant/go-gpui"
)

//go:embed input.html
var inputHTML string

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 420
	DefaultHeight = 560
)

// App is the input screen.
type App struct {
	page *gpui.Page
}

// New parses the embedded template.
func New() (*App, error) {
	p, err := gpui.New(gpui.Config{
		Title:  "Input",
		HTML:   inputHTML,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	app := &App{page: p}
	p.Handle(gpui.Handlers{Click: app.click})

	return app, nil
}

// click scrolls to the end when the link is pressed.
func (a *App) click(_ context.Context, box gpui.Box) error {
	if box.ID == "more" {
		a.scrollTo(0, 1<<30)
	}

	return nil
}

// scroller is the page scroll surface the window consumes. The assertion
// keeps the example building before the page side of v0.0.2 lands.
type scroller interface {
	ScrollTo(x, y int)
}

func (a *App) scrollTo(x, y int) {
	if sc, ok := any(a.page).(scroller); ok {
		sc.ScrollTo(x, y)
	}
}
