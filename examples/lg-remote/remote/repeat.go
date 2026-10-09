package remote

import (
	"sync/atomic"
	"time"
)

type repeatState struct {
	id, host, mode string
	start, next    time.Time
	epoch          atomic.Uint64
	pending        atomic.Bool
}

func repeatable(id string) bool {
	switch id {
	case "volup", "voldn", "chup", "chdn", "up", "down", "left", "right":
		return true
	}
	return false
}

func (a *App) armRepeat(id string) {
	a.stopRepeat()
	if !repeatable(id) {
		return
	}
	r := &a.repeat
	r.id, r.host, r.mode = id, a.hostValue(), a.view.Mode
	r.start = time.Now()
	r.next = r.start.Add(350 * time.Millisecond)
}

func (a *App) stopRepeat() {
	a.repeat.id = ""
	a.repeat.epoch.Add(1)
}

func (a *App) repeatTick() bool {
	r := &a.repeat
	if r.id == "" {
		return false
	}
	if (a.page.PressedID() != r.id && a.nativeHeld != r.id) || a.view.Mode != r.mode || a.hostValue() != r.host {
		a.stopRepeat()
		return false
	}
	now := time.Now()
	if now.Before(r.next) {
		return false
	}
	interval := 120 * time.Millisecond
	if now.Sub(r.start) >= 2*time.Second {
		interval = 80 * time.Millisecond
	}
	if now.Sub(r.start) >= 4*time.Second {
		interval = 50 * time.Millisecond
	}
	r.next = now.Add(interval)
	if r.mode == "Bluetooth" {
		a.sendBT(r.id)
		return true
	}
	if !r.pending.CompareAndSwap(false, true) {
		return false
	}
	host, epoch := r.host, r.epoch.Load()
	spec, _ := actionOf(r.id)
	if !a.later(func() {
		defer r.pending.Store(false)
		if r.epoch.Load() == epoch {
			a.doControl(host, spec, true)
		}
	}) {
		r.pending.Store(false)
	}
	return false
}
