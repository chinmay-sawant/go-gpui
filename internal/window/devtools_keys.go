package window

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// devKeys are the keys the overlay owns while a page has an inspector.
var devKeys = [...]ebiten.Key{ebiten.KeyF12, ebiten.KeyI, ebiten.KeyO}

// drawTimer is the optional hook a screen implements to receive the window's
// draw time. host.Inspector keeps its three-method shape.
type drawTimer interface {
	SetDrawTime(time.Duration)
}

// devtoolsKeys steps the overlay's keys before the page sees them. A key the
// overlay owns is swallowed until it goes up, and a swallowed key also drops
// the printable input of the frame, so "o" toggles the operation view
// without typing.
func (s *shell) devtoolsKeys(mods modifiers) (bool, error) {
	ate := false

	for _, key := range devKeys {
		swallowed, err := s.devKeyStep(key, ebiten.IsKeyPressed(key), mods)
		if err != nil {
			return false, err
		}

		ate = ate || swallowed
	}

	return ate, nil
}

// devKeyStep advances the overlay watch for one key and reports whether the
// key belongs to the overlay this frame.
func (s *shell) devKeyStep(key ebiten.Key, pressed bool, mods modifiers) (bool, error) {
	i := devKeyIndex(key)
	if i < 0 {
		return false, nil
	}

	if !s.dev.eaten[i] && !s.devOwns(key, mods) {
		return false, nil
	}

	down, _ := s.dev.watch.step(key, pressed)
	if !pressed {
		swallowed := s.dev.eaten[i]
		s.dev.eaten[i] = false

		return swallowed, nil
	}

	s.dev.eaten[i] = true
	if !down {
		return true, nil
	}

	return true, s.devKeyFired(key)
}

// devKeyFired runs the action of one overlay key.
func (s *shell) devKeyFired(key ebiten.Key) error {
	if key == ebiten.KeyO {
		s.dev.ops = !s.dev.ops
		s.dev.tab = devTabOps

		return nil
	}

	return s.devToggle()
}
