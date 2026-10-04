package devtools

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

// New parses the embedded template, registers the counter, and hands the
// template an image through SetImage. DevTools starts the overlay on.
func New() (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:    "DevTools",
		HTML:     devtoolsHTML,
		Width:    DefaultWidth,
		Height:   DefaultHeight,
		DevTools: true,
	})
	if err != nil {
		return nil, err
	}

	app := &App{page: page}
	page.SetImage("shot", shotPNG())
	page.Handle(gpui.Handlers{Click: app.onClick})
	page.SetData(&app.view)

	return app, nil
}

// onClick counts the counter increments.
func (a *App) onClick(_ context.Context, box gpui.Box) error {
	if box.ID != "counter" {
		return nil
	}

	a.view.Count++
	a.page.SetData(&a.view)

	return nil
}
