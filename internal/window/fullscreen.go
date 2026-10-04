package window

import (
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

// fullscreenSupported reports that the platform owns a window that can go
// fullscreen. wasm and mobile keep their own fullscreen.
func fullscreenSupported() bool {
	switch runtime.GOOS {
	case "js", "android", "ios":
		return false
	}

	return true
}

// fullscreenNow reports the current fullscreen state.
func (s *shell) fullscreenNow() bool {
	if s.readFullscreen != nil {
		return s.readFullscreen()
	}

	if !fullscreenSupported() {
		return false
	}

	return ebiten.IsFullscreen()
}

// setFullscreen changes the fullscreen state.
func (s *shell) setFullscreen(on bool) {
	if s.applyFullscreen != nil {
		s.applyFullscreen(on)

		return
	}

	if fullscreenSupported() {
		ebiten.SetFullscreen(on)
	}
}

// f11 toggles fullscreen and reports whether the window consumed the key.
func (s *shell) f11() bool {
	if !fullscreenSupported() {
		return false
	}

	s.setFullscreen(!s.fullscreenNow())

	return true
}
