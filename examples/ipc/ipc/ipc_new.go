package ipc

import (
	"github.com/chinmay-sawant/go-gpui"
)

// New parses the embedded template, registers the demo.log listeners and the
// demo.double handler, and wires the click handler.
func New() (*App, error) {
	page, err := gpui.New(gpui.Config{
		Title:  "IPC",
		HTML:   ipcHTML,
		Width:  DefaultWidth,
		Height: DefaultHeight,
	})
	if err != nil {
		return nil, err
	}

	app := &App{page: page}
	app.register()
	app.view.Status = "Ready: click a control above and the outcome appears here."
	page.Handle(gpui.Handlers{Click: app.onClick})
	page.SetData(&app.view)

	return app, nil
}

// register adds two listeners on demo.log and one handler on demo.double.
// Cancel removes them, and a later register adds fresh ones.
// Two listeners show that Send fans out; one handler shows that Request asks
// a single callback for a reply.
func (a *App) register() {
	a.stopLog = gpui.Listen("demo.log", a.onLog)
	a.stopCount = gpui.Listen("demo.log", a.onCount)
	a.stopDouble = gpui.Handle("demo.double", a.onDouble)
	a.view.Live = true
}
