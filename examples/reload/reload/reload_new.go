package reload

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

// New reads index.html and installs the counter click handler. reload false
// turns the file watch off.
func New(reload bool) (*App, error) {
	return NewFile(SourcePath, reload)
}

// NewFile reads path as the page source, so a test can point a temp copy at
// the same template.
func NewFile(path string, reload bool) (*App, error) {
	p, err := gpui.New(gpui.Config{
		Title:            "Hot reload",
		File:             path,
		DisableHotReload: !reload,
		Width:            DefaultWidth,
		Height:           DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	return finish(p)
}

// NewHTML builds the same page from a string, for tests.
func NewHTML(html string, reload bool) (*App, error) {
	p, err := gpui.New(gpui.Config{
		Title:            "Hot reload",
		HTML:             html,
		DisableHotReload: !reload,
		Width:            DefaultWidth,
		Height:           DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	return finish(p)
}

func finish(p *gpui.Page) (*App, error) {
	app := &App{page: p}
	p.Handle(gpui.Handlers{Click: app.onClick})
	p.SetData(&app.view)

	return app, nil
}

// onClick counts clicks on the bump button. Click draws the page after the
// handler returns.
func (a *App) onClick(_ context.Context, box gpui.Box) error {
	if box.ID != "bump" {
		return nil
	}

	a.view.Count++

	return nil
}
