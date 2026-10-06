package dino

import (
	"context"
	"math"
	"time"

	"github.com/chinmay-sawant/go-gpui"
)

// duckHold is how long a swipe down keeps the crouch.
const duckHold = 600 * time.Millisecond

// BindTouch replaces the key-only handlers with the phone input: a tap or
// a swipe up jumps, a swipe down ducks, and either starts the first run or
// starts a new run after a crash.
func (a *App) BindTouch() {
	a.touch = true

	a.page.Handle(gpui.Handlers{
		KeyDown: a.onKeyDown,
		KeyUp:   a.onKeyUp,
		Click:   a.onTap,
		Swipe:   a.onSwipe,
	})
}

// onTap jumps like a key press.
func (a *App) onTap(_ context.Context, _ gpui.Box) error {
	a.jumpPress()

	return nil
}

// onSwipe maps a one-finger swipe to the game: up jumps, down ducks for
// duckHold, and a sideways swipe is ignored.
func (a *App) onSwipe(_ context.Context, dx, dy float64) error {
	switch {
	case dy < 0 && -dy > math.Abs(dx):
		a.jumpPress()
	case dy > 0 && dy > math.Abs(dx):
		a.game.setDuck(true)
		a.duckUntil = a.now().Add(duckHold)
	}

	return nil
}

// jumpPress is the tap's press logic: it cancels any duck, then jumps like
// a key press. A touch jump has no release, so the handler keeps the jump
// held while the dinosaur rises: every tap reaches the high jump.
func (a *App) jumpPress() {
	a.duckUntil = time.Time{}
	a.game.setDuck(false)

	ground := a.game.onGround

	a.game.jump()

	if ground && !a.game.onGround {
		a.tapped = true
		a.game.jumpHeld = true
	}
}

// duckStep releases a swipe duck once its hold ends.
func (a *App) duckStep(now time.Time) {
	if !a.duckUntil.IsZero() && !now.Before(a.duckUntil) {
		a.game.setDuck(false)
		a.duckUntil = time.Time{}
	}
}

// tapStep releases a tap jump once the dinosaur stops rising, or when the
// run ends, so the tap neither floats nor bounces on landing.
func (a *App) tapStep() {
	if !a.tapped {
		return
	}

	if a.game.onGround || a.game.phase != running || a.game.vy <= 0 {
		a.game.jumpHeld = false
		a.tapped = false
	}
}
