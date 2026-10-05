// Package app is the Microsoft Teams app shell. One HTML template holds the
// app rail and the menu fragment each menu package provides. View carries
// every menu's data, and a rail click switches menus. gpui opens the window;
// this package does not.
package app

import (
	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/teams/store"
)

// App is the Teams app screen.
type App struct {
	page  *gpui.Page
	view  View
	store *store.Store
}

// New parses the app template, opens the state database, registers its
// images and handlers. The stored theme chooses the stylesheet, so a saved
// light state starts light and the profile toggle still works.
func New(opts ...Option) (*App, error) {
	cfg := config{dbPath: store.Memory}

	for _, opt := range opts {
		opt(&cfg)
	}

	st := openStore(cfg.dbPath)
	app := &App{store: st}
	app.view = app.loadView(st)

	page, err := gpui.New(gpui.Config{
		Title:     "Microsoft Teams",
		HTML:      buildHTML(),
		Theme:     themeSource(app.view.Dark),
		Width:     DefaultWidth,
		Height:    DefaultHeight,
		MinWidth:  MinWidth,
		MinHeight: MinHeight,
	})
	if err != nil {
		return nil, err
	}

	app.page = page
	registerImages(page, app.view.Dark)
	page.Handle(gpui.Handlers{Click: app.onClick, Change: app.onChange, Submit: app.onSubmit})
	page.SetData(app.view)

	return app, nil
}
