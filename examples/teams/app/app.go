// Package app is the Microsoft Teams app shell. One HTML template holds the
// app rail and the menu fragment each menu package provides. View carries
// every menu's data, and a rail click switches menus. gpui opens the window;
// this package does not.
package app

import (
	"github.com/chinmay-sawant/go-gpui"
	"github.com/chinmay-sawant/go-gpui/examples/teams/store"
)

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 1400
	DefaultHeight = 900

	// MinWidth and MinHeight are the smallest frame the screen will draw.
	MinWidth  = 820
	MinHeight = 600
)

// Option configures New.
type Option func(*config)

// config holds the settings an option can change.
type config struct {
	dbPath string
}

// WithDB keeps the state in the SQLite database at path. The default is an
// in-memory database, so a caller that passes no option never touches disk.
func WithDB(path string) Option {
	if path == "" {
		path = store.Memory
	}

	return func(c *config) {
		c.dbPath = path
	}
}

// App is the Teams app screen.
type App struct {
	page  *gpui.Page
	view  View
	store *store.Store
}

// New parses the app template, opens the state database, registers its
// images and handlers.
func New(opts ...Option) (*App, error) {
	cfg := config{dbPath: store.Memory}

	for _, opt := range opts {
		opt(&cfg)
	}

	page, err := gpui.New(gpui.Config{
		Title:     "Microsoft Teams",
		HTML:      buildHTML(),
		Theme:     themeSource(true),
		Width:     DefaultWidth,
		Height:    DefaultHeight,
		MinWidth:  MinWidth,
		MinHeight: MinHeight,
	})
	if err != nil {
		return nil, err
	}

	st := openStore(cfg.dbPath)
	app := &App{page: page, store: st}
	app.view = app.loadView(st)
	registerImages(page, app.view.Dark)
	page.Handle(gpui.Handlers{Click: app.onClick, Change: app.onChange, Submit: app.onSubmit})
	page.SetData(app.view)

	return app, nil
}
