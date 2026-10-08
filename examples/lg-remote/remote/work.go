package remote

func (a *App) doExec(host, spec string) {
	msg, err := a.use().Exec(host, spec)
	if err != nil {
		a.note(err.Error(), "", "")
		return
	}

	a.note(msg, a.use().SavedModel(), "")
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
		a.note(err.Error(), name, host)
		return
	}

	a.note(msg, a.use().SavedModel(), host)
}

func (a *App) doWake() {
	msg, err := a.use().Wake()
	if err != nil {
		a.note(err.Error(), "", "")
		return
	}

	a.note(msg, "", "")
}

func (a *App) note(status, title, host string) {
	select {
	case a.notes <- update{status: status, title: title, host: host}:
	default:
	}
}
