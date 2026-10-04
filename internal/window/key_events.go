package window

import "github.com/hajimehoshi/ebiten/v2"

// keyEvents sends each key press and release to the screen in the frame it
// happens. A key handler does not draw; the screen paints on its own.
// The window handles Tab and Escape before the screen sees them.
func (s *shell) keyEvents(mods modifiers) error {
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
			taken, err := s.windowKey(key, true, mods)
			if err != nil {
				return err
			}

			if !taken {
				if err := s.app.KeyDown(s.ctx, name); err != nil {
					return err
				}
			}
		}

		if up {
			taken, err := s.windowKey(key, false, mods)
			if err != nil {
				return err
			}

			if !taken {
				if err := s.app.KeyUp(s.ctx, name); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
