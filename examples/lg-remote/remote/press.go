package remote

import "github.com/chinmay-sawant/ownframe/examples/lg-remote/bridge"

func (a *App) toggleMode() {
	if !a.view.ShowBluetooth {
		return
	}

	if a.view.Mode == "Bluetooth" {
		bridge.Enqueue("stop")
		a.view.Mode = "Wi-Fi"
		a.view.Hint = "Same Wi-Fi as the TV."
		a.view.Status = "Wi-Fi"
		return
	}

	a.view.Mode = "Bluetooth"
	a.btSeen = 0
	a.view.Hint = "Pair this phone in the TV Bluetooth menu. Dimmed controls need Wi-Fi. Wake uses Wi-Fi."
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

	if id == "wake" {
		a.changePower(true)
		return
	}
	if id == "power" && a.view.Mode == "Bluetooth" {
		if a.sendBT(id) {
			a.powerIntent.seq++
			a.powerIntent.pending = false
			a.view.PowerOn = !a.view.PowerOn
		}
		return
	}
	if id == "power" && a.view.Mode == "Wi-Fi" {
		a.changePower(!a.view.PowerOn)
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
	a.later(func() { a.doControl(host, spec, false) })
}

func (a *App) sendBT(id string) bool {
	cmd, ok := btCommand(id)
	if !ok {
		a.view.Status = "That control needs Wi-Fi"
		return false
	}

	if !bridge.Enqueue(cmd) {
		a.view.Status = "Bluetooth queue is full. Try again."
		return false
	}
	a.view.Status = "Queued for Bluetooth"
	return true
}
