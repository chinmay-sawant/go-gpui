package wire

import "github.com/chinmay-sawant/ownframe/examples/download-manager/ui"

// run executes one command on the worker goroutine.
func (b *Backend) run(c command) {
	switch c.kind {
	case cmdAdd:
		b.runAdd(c.add)
	case cmdControl:
		b.runControl(c.ctl)
	case cmdActive:
		b.pushActive()
	case cmdPage:
		b.runPage(c.page)
	case cmdSummary:
		b.runSummary(c.gen)
	case cmdDark:
		if err := writeDark(b.themePath(), c.dark); err != nil {
			b.notice("Theme not saved: " + err.Error())
		}
	}
}

// pushActive sends the current active set. It is the only path that reads
// the engine's live list, so the UI loop never touches engine memory.
func (b *Backend) pushActive() {
	b.push(ui.Update{Kind: ui.UpdateActive, Active: b.activeRows()})
}
