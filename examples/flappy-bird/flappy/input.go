package flappy

import (
	"context"

	"github.com/chinmay-sawant/go-gpui"
)

// onKeyDown maps a key press to a flap or a restart. The handler never
// draws; the tick paints the result.
func (a *App) onKeyDown(_ context.Context, key string) error {
	switch key {
	case "space", "arrowup", "w", "enter":
		a.game.flap()
	case "r":
		if a.game.phase == over {
			a.game.again()
		}
	}

	return nil
}

// onClick flaps from a pointer press anywhere in the picture, and restarts
// after a crash, as the flap key does.
func (a *App) onClick(_ context.Context, _ gpui.Box) error {
	a.game.flap()

	return nil
}
