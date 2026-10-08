package remote

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/lg-remote/bridge"
)

func (a *App) onTick(ctx context.Context) error {
	changed := a.load()
	if a.drain() {
		changed = true
	}
	if radio := bridge.Bluetooth(); radio != "" && radio != a.btSeen {
		a.btSeen = radio
		a.view.Status = radio
		changed = true
	}

	if !changed {
		return nil
	}

	a.page.SetData(&a.view)

	return a.page.Redraw(ctx)
}

func (a *App) load() bool {
	if a.loaded {
		return false
	}

	a.loaded = true
	link := a.use()
	if host := link.SavedHost(); host != "" {
		a.page.SetFormValue("host", host)
	}

	dirty := false
	if model := link.SavedModel(); model != "" {
		a.view.Title = model
		dirty = true
	}

	if !a.wantSearch {
		return dirty
	}

	a.view.Status = "Searching this Wi-Fi"
	a.later(a.doFind)

	return true
}

func (a *App) drain() bool {
	changed := false

	for {
		select {
		case u := <-a.notes:
			if u.status != "" {
				a.view.Status = u.status
			}
			if u.title != "" {
				a.view.Title = u.title
			}
			if u.host != "" {
				a.page.SetFormValue("host", u.host)
			}
			if u.setPower {
				a.view.PowerOn = u.powerOn
			}
			changed = true
		default:
			return changed
		}
	}
}
