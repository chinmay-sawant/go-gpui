package window

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/go-gpui/internal/host"
)

func (s *shell) Update() error {
	if s.app == nil {
		return errNilApp
	}

	if err := s.ctx.Err(); err != nil {
		return err
	}

	if err := s.devSync(); err != nil {
		return err
	}

	if err := s.keys(); err != nil {
		return err
	}

	if err := s.resize(); err != nil {
		return err
	}

	s.updatePassthrough()
	if err := s.pointer(); err != nil {
		return err
	}

	if err := s.dropPass(ebiten.DroppedFiles()); err != nil {
		return err
	}

	if err := s.tickFrame(); err != nil {
		return err
	}

	if err := s.pollReload(); err != nil {
		return err
	}

	if err := s.syncImage(); err != nil {
		return err
	}

	s.applyScrollRequest()
	s.devRefresh()
	s.wheel()

	return nil
}

// tickFrame runs the screen's per-frame callback when it has one.
func (s *shell) tickFrame() error {
	ticker, ok := s.app.(host.Ticker)
	if !ok {
		return nil
	}

	return ticker.Tick(s.ctx)
}
