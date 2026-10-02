package ipc

import (
	"context"
	"strconv"

	"github.com/chinmay-sawant/go-gpui"
)

// New parses the embedded template and registers the demo.log listener, the
// demo.double handler, and the click handler.
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
	app.stopLog = gpui.Listen("demo.log", app.onLog)
	app.stopDouble = gpui.Handle("demo.double", app.onDouble)
	page.Handle(gpui.Handlers{Click: app.onClick})
	page.SetData(&app.view)

	return app, nil
}

// onLog appends each payload to the log the template prints.
func (a *App) onLog(payload string) {
	if a.view.Log != "" {
		a.view.Log += " "
	}

	a.view.Log += payload
}

// onDouble parses the payload and returns twice its value.
func (a *App) onDouble(_ context.Context, payload string) (string, error) {
	n, err := strconv.Atoi(payload)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(n * 2), nil
}
