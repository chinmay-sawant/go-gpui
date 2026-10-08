package remote

import (
	"context"
	"fmt"
)

func (a *App) onSwipe(ctx context.Context, dx, dy float64) error {
	if a.view.Panel != "pad" {
		return nil
	}

	if a.view.Mode == "Bluetooth" {
		a.view.Status = "The pad needs Wi-Fi"
		a.page.SetData(&a.view)
		return a.page.Redraw(ctx)
	}

	x := clampInt(int(dx), -160, 160)
	y := clampInt(int(dy), -160, 160)
	if x == 0 && y == 0 {
		return nil
	}

	host := a.page.FormValue("host")
	spec := fmt.Sprintf("move:%d,%d", x, y)
	a.later(func() { a.doExec(host, spec) })

	return nil
}

func (a *App) padPress(id string) bool {
	spec := ""
	switch id {
	case "pad", "clicker":
		spec = "click:"
	case "wheelup":
		spec = "scroll:-3"
	case "wheeldown":
		spec = "scroll:3"
	default:
		return false
	}

	if a.view.Mode == "Bluetooth" {
		a.view.Status = "The pad needs Wi-Fi"
		return true
	}

	host := a.host
	a.later(func() { a.doExec(host, spec) })

	return true
}

func clampInt(v, lo, hi int) int {
	if v < lo {
		return lo
	}

	if v > hi {
		return hi
	}

	return v
}
