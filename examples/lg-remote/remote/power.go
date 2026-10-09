package remote

type powerIntent struct {
	seq     uint64
	pending bool
}

func (a *App) changePower(on bool) {
	previous := a.view.PowerOn
	a.powerIntent.seq++
	a.powerIntent.pending = true
	seq := a.powerIntent.seq
	a.view.PowerOn = on
	host := a.host
	a.view.Status = "Turning the TV off"
	if on {
		a.view.Status = "Waking the TV"
	}
	if !a.later(func() { a.doPower(host, on, previous, seq) }) {
		a.view.PowerOn = previous
		a.powerIntent.pending = false
	}
}

func (a *App) doPower(host string, on, previous bool, seq uint64) {
	var msg string
	var err error
	if on {
		msg, err = a.use().Wake()
	} else {
		msg, err = a.use().Exec(host, "power:")
	}
	if err != nil {
		if !on && offline(err) {
			previous = on
			msg = "Power requested. The TV connection closed."
		} else {
			msg = err.Error()
		}
		a.notes <- update{status: msg, setPower: true, powerOn: previous, powerSeq: seq}
		return
	}
	a.notes <- update{status: msg, setPower: true, powerOn: on, powerSeq: seq}
	if on {
		a.noteWake(a.use().SavedHost())
	}
}
