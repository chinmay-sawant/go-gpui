package wire

import (
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// runControl applies one job command and refreshes the active set.
func (b *Backend) runControl(c ui.Control) {
	var err error

	switch c.Action {
	case ui.ControlPause:
		err = b.eng.Pause(b.ctx, c.ID)
	case ui.ControlResume:
		err = b.eng.Resume(b.ctx, c.ID)
	case ui.ControlCancel:
		err = b.eng.Cancel(b.ctx, c.ID)
	case ui.ControlRetry:
		err = b.eng.Retry(b.ctx, c.ID)
	case ui.ControlRemove:
		err = b.store.DeleteJob(b.ctx, c.ID)
	}

	if err != nil {
		b.notice("Command failed: " + err.Error())

		return
	}

	b.pushActive()
}
