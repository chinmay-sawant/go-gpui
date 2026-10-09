package remote

import (
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/lg-remote/bridge"
)

func (a *App) doExec(host, spec string) {
	msg, err := a.use().Exec(host, spec)
	if err != nil {
		if offline(err) {
			a.noteState(err.Error(), "", "", false)
			if spec == "power:" {
				a.doWake()
			}
			return
		}

		a.note(err.Error(), "", "")
		return
	}

	on := spec != "power:"
	if spec == "power:" {
		msg = "Turning the TV off"
	}

	a.noteState(msg, a.use().SavedModel(), "", on)
}

func offline(err error) bool {
	if err == nil {
		return false
	}

	text := err.Error()

	return strings.Contains(text, "did not answer") ||
		strings.Contains(text, "closed the connection")
}

func (a *App) doFind() {
	host, err := a.use().Scan()
	if err != nil {
		a.note(err.Error(), "", "")
		return
	}

	name := a.use().SavedModel()
	label := host
	if name != "" {
		label = name
	}

	a.note("Found "+label, name, host)
	msg, err := a.use().Exec(host, "pair:")
	if err != nil {
		if offline(err) {
			a.noteState(err.Error(), name, host, false)
			return
		}

		a.note(err.Error(), name, host)
		return
	}

	a.noteState(msg, a.use().SavedModel(), host, true)
}

func (a *App) doWake() {
	msg, err := a.use().Wake()
	if err != nil {
		a.note(err.Error(), "", "")
		return
	}

	a.noteState(msg, "", "", false)
	if a.phone {
		bridge.Enqueue("con:48")
	}
	a.noteWake(a.use().SavedHost())
}

func (a *App) note(status, title, host string) {
	a.notes <- update{status: status, title: title, host: host}
}

func (a *App) noteState(status, title, host string, on bool) {
	a.notes <- update{status: status, title: title, host: host, setPower: true, powerOn: on}
}
