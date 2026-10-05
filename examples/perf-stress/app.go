package main

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/go-gpui"
)

// App is the stress dashboard screen.
type App struct {
	page  *gpui.Page
	view  View
	rows  int
	state tickState
	track gpui.Box
	eq    gpui.Box
}

// newApp parses the dashboard template and wires input and the tick.
func newApp(rows int) (*App, error) {
	page, err := gpui.NewWithOptions(gpui.Config{
		Title: "Stress", HTML: buildHTML(),
		Width: 1280, Height: 900, MinWidth: 640, MinHeight: 480,
	}, gpui.WithPerf(true))
	if err != nil {
		return nil, err
	}
	a := &App{page: page, view: makeView(rows), rows: rows}
	page.Handle(gpui.Handlers{Click: a.onClick, Change: a.onChange})
	page.SetTick(a.Tick)
	page.SetData(a.view)
	return a, nil
}

// onClick switches the sidebar section or bumps progress, then redraws.
func (a *App) onClick(_ context.Context, box gpui.Box) error {
	if strings.HasPrefix(box.ID, "nav-") {
		a.view.Active = strings.TrimPrefix(box.ID, "nav-")
	}
	if box.ID == "go" {
		a.view.Progress = (a.view.Progress + 10) % 101
	}
	if box.ID == "reset" {
		a.view = makeView(a.rows)
	}
	a.page.SetData(a.view)
	return nil
}

// onChange keeps the filter value by re-setting the view.
func (a *App) onChange(_ context.Context, _ gpui.Box) error {
	a.page.SetData(a.view)
	return nil
}
