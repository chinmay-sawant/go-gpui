package flappy

import "github.com/chinmay-sawant/go-gpui"

// paintPipes places the four fills of every live pipe and hides the rest.
func (a *App) paintPipes(d *gpui.Display) {
	for slot := range pipeSlots {
		if slot < len(a.game.pipes) {
			a.paintPipe(d, slot, a.game.pipes[slot])

			continue
		}

		for i := range pipePartCount {
			setFill(d, a.parts.pipes[slot][i], rect{}, false)
		}
	}
}

// paintPipe places the four fills of the pair in one slot.
func (a *App) paintPipe(d *gpui.Display, slot int, p pipe) {
	parts := a.game.pipeRects(p)

	for i := range pipePartCount {
		setFill(d, a.parts.pipes[slot][i], parts[i], true)
	}
}
