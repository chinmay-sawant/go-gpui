package dino

import "context"

// onKeyDown maps a key press to a game action. The handler never draws;
// the tick paints the result.
func (a *App) onKeyDown(_ context.Context, key string) error {
	switch key {
	case "space", "arrowup", "w", "enter":
		a.game.jumpHeld = true
		a.game.jump()
	case "arrowdown", "s":
		a.game.setDuck(true)
	case "r":
		if a.game.phase == over {
			a.game.again()
		}
	}

	return nil
}

// onKeyUp releases the jump or the crouch.
func (a *App) onKeyUp(_ context.Context, key string) error {
	switch key {
	case "space", "arrowup", "w", "enter":
		a.game.jumpHeld = false
	case "arrowdown", "s":
		a.game.setDuck(false)
	}

	return nil
}
