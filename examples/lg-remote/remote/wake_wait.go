package remote

import "time"

func (a *App) waitForTV() {
	if !a.async {
		return
	}

	host := a.host
	if host == "" {
		host = a.use().SavedHost()
	}

	for _, pause := range []time.Duration{8 * time.Second, 10 * time.Second} {
		time.Sleep(pause)
		msg, err := a.use().Exec(host, "pair:")
		if err != nil {
			continue
		}

		a.noteState(msg, a.use().SavedModel(), "", true)

		return
	}

	a.noteState("Wake sent. The TV is still starting.", "", "", false)
}
