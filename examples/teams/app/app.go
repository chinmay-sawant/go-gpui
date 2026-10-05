// Package app is the Microsoft Teams app shell. One HTML template holds the
// app rail and the menu fragment each menu package provides. View carries
// every menu's data, and a rail click switches menus. gpui opens the window;
// this package does not.
package app

import "github.com/chinmay-sawant/go-gpui"

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 1400
	DefaultHeight = 900

	// MinWidth and MinHeight are the smallest frame the screen will draw.
	MinWidth  = 820
	MinHeight = 600
)

// App is the Teams app screen.
type App struct {
	page *gpui.Page
	view View
}

// New parses the app template, registers its images and handlers.
func New() (*App, error) {
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

	app := &App{page: page, view: DefaultView()}
	registerImages(page, true)
	page.Handle(gpui.Handlers{Click: app.onClick, Change: app.onChange, Submit: app.onSubmit})
	page.SetData(app.view)

	return app, nil
}
