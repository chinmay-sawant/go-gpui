package remote

import (
	"errors"
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/lg-remote/tv"
)

type stepStarter interface {
	StartStep(host, spec string) (func() (string, error), error)
}

func (a *App) doControl(host, spec string, repeat bool) {
	link := a.use()
	start, ok := link.(stepStarter)
	if !ok || (!strings.HasPrefix(spec, "vol:") && !strings.HasPrefix(spec, "ch:")) {
		a.doExec(host, spec)
		return
	}
	wait, err := start.StartStep(host, spec)
	if err != nil {
		if !repeat || !errors.Is(err, tv.ErrStepsBusy) {
			a.note(err.Error(), "", "")
		}
		return
	}
	finish := func() {
		msg, err := wait()
		if err != nil {
			a.note(err.Error(), "", "")
			return
		}
		a.note(msg, "", "")
	}
	if a.async {
		go finish()
	} else {
		finish()
	}
}
