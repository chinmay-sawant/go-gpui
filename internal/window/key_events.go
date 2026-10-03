package window

import "github.com/hajimehoshi/ebiten/v2"

// keyEvents sends each key press and release to the screen in the frame it
// happens. A key handler does not draw; the screen paints on its own.
func (s *shell) keyEvents() error {
	for key := ebiten.Key(0); key <= ebiten.KeyMax; key++ {
		down, up := s.watched.step(key, ebiten.IsKeyPressed(key))
		if !down && !up {
			continue
		}

		name := keyName(key)
		if name == "" {
			continue
		}

		if down {
			if err := s.app.KeyDown(s.ctx, name); err != nil {
				return err
			}
		}

		if up {
			if err := s.app.KeyUp(s.ctx, name); err != nil {
				return err
			}
		}
	}

	return nil
}
