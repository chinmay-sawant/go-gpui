package remote

import "github.com/chinmay-sawant/ownframe/examples/lg-remote/bridge"

func (a *App) toggleMode() {
	if !a.view.ShowBluetooth {
		return
	}

	if a.view.Mode == "Bluetooth" {
		a.view.Mode = "Wi-Fi"
		a.view.Hint = "Same Wi-Fi as the TV."
		a.view.Status = "Wi-Fi"
		return
	}

	a.view.Mode = "Bluetooth"
	a.view.Hint = "Pair this phone in the TV Bluetooth menu. Wake still uses Wi-Fi."
	a.view.Status = "Bluetooth"
	bridge.Enqueue("start")
}

func (a *App) connect() {
	if a.view.Mode == "Bluetooth" {
		bridge.Enqueue("start")
		a.view.Status = "On the TV, open Bluetooth and pair this phone."
		return
	}

	a.view.Status = "Connecting. Accept the prompt on the TV if it appears."
	host := a.host
	a.later(func() { a.doExec(host, "pair:") })
}

func (a *App) press(id string) {
	if a.padPress(id) {
		return
	}

	if id == "power" && !a.view.PowerOn {
		a.view.Status = "Waking the TV"
		a.later(a.doWake)
		return
	}

	if id == "wake" {
		a.view.Status = "Sending wake"
		a.later(a.doWake)
		return
	}

	if !a.view.PowerOn && a.view.Mode != "Bluetooth" {
		a.view.Status = "The TV is off. Press the green Power key."
		return
	}

	if a.view.Mode == "Bluetooth" {
		a.sendBT(id)
		return
	}

	spec, ok := actionOf(id)
	if !ok {
		return
	}

	host := a.host
	a.view.Status = "Sending"
	a.later(func() { a.doExec(host, spec) })
}

func (a *App) sendBT(id string) {
	cmd, ok := btCommand(id)
	if !ok {
		a.view.Status = "That control needs Wi-Fi"
		return
	}

	bridge.Enqueue(cmd)
	a.view.Status = "Sent"
}

func (a *App) later(fn func()) {
	if a.async {
		go fn()
		return
	}

	fn()
}
