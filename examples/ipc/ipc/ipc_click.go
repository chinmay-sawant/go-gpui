package ipc

import (
	"context"
	"errors"

	"github.com/chinmay-sawant/go-gpui"
)

// onClick runs the action for the clicked button.
func (a *App) onClick(ctx context.Context, box gpui.Box) error {
	switch box.ID {
	case "send":
		gpui.Send("demo.log", "ping")
		a.view.Status = "sent ping"
	case "request":
		reply, err := gpui.Request(ctx, "demo.double", a.page.FormValue("num"))
		if err != nil {
			a.view.Status = err.Error()
			return nil
		}

		a.view.Status = reply
	case "missing":
		_, err := gpui.Request(ctx, "demo.none", "x")
		if errors.Is(err, gpui.ErrNoHandler) {
			a.view.Status = "demo.none: " + err.Error()
			break
		}

		a.view.Status = "demo.none: unexpected reply"
	case "cancel":
		a.Cancel()
		gpui.Send("demo.log", "after-cancel")
		a.view.Status = "listener canceled; send ignored"
	}

	return nil
}
