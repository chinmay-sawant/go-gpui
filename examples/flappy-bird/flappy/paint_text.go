package flappy

import (
	"strconv"

	"github.com/chinmay-sawant/go-gpui"
)

// paintText refreshes the score and shows the ready or the game-over text.
func (a *App) paintText(d *gpui.Display) {
	isReady := a.game.phase == ready
	isOver := a.game.phase == over

	reveal(a.parts.title, isReady, "FLAPPY BIRD")
	reveal(a.parts.hint, isReady, "SPACE OR CLICK TO FLAP")
	reveal(a.parts.over, isOver, "GAME OVER")
	reveal(a.parts.oscore, isOver, "SCORE "+strconv.Itoa(a.game.score))
	reveal(a.parts.obest, isOver, "BEST "+strconv.Itoa(a.game.best))
	reveal(a.parts.again, isOver, "PRESS R OR SPACE")

	setText(a.parts.score, scoreText(a.game.score, isOver))
	setFill(d, a.parts.board, a.parts.boardRect, isOver)
}

// scoreText is the running score, hidden once the game is over.
func scoreText(score int, over bool) string {
	if over {
		return ""
	}

	return strconv.Itoa(score)
}

// reveal writes a message, or empties the run that does not apply.
func reveal(op *gpui.DisplayOp, visible bool, text string) {
	if !visible {
		text = ""
	}

	setText(op, text)
}

// setText assigns a text run. An empty string hides it without removing the
// element from the page.
func setText(op *gpui.DisplayOp, text string) {
	if op != nil && op.Text != text {
		op.Text = text
	}
}
