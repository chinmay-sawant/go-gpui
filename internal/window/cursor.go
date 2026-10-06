package window

import (
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/ownframe/internal/host"
)

// cursorShape maps a page shape to an Ebiten cursor shape.
func cursorShape(shape host.Shape) ebiten.CursorShapeType {
	switch shape {
	case host.ShapeText:
		return ebiten.CursorShapeText
	case host.ShapePointer:
		return ebiten.CursorShapePointer
	case host.ShapeResizeEW:
		return ebiten.CursorShapeEWResize
	case host.ShapeResizeNS:
		return ebiten.CursorShapeNSResize
	default:
		return ebiten.CursorShapeDefault
	}
}

// cursorSupported reports that the platform has a mouse cursor to shape.
// wasm and mobile keep their own cursor.
func cursorSupported() bool {
	switch runtime.GOOS {
	case "js", "android", "ios":
		return false
	}

	return true
}

// applyCursor reads the hovered shape and sets the cursor only when it
// changed. The last shape stays on the shell for tests.
func (s *shell) applyCursor() {
	shaper, ok := s.app.(host.CursorShape)
	if !ok {
		return
	}

	s.setCursorShape(cursorShape(shaper.CursorShape()))
}

// applyThumbCursor shows the resize cursor over a scrollbar thumb. The page
// does not see the thumbs, so the shell names these shapes itself.
func (s *shell) applyThumbCursor(axis int) {
	switch axis {
	case dragVertical:
		s.setCursorShape(ebiten.CursorShapeNSResize)
	case dragHorizontal:
		s.setCursorShape(ebiten.CursorShapeEWResize)
	}
}

// setCursorShape records and applies one cursor shape, calling the test hook
// or Ebiten only when the shape changed.
func (s *shell) setCursorShape(next ebiten.CursorShapeType) {
	if next == s.cursor {
		return
	}

	s.cursor = next

	if s.setCursor != nil {
		s.setCursor(next)

		return
	}

	if cursorSupported() {
		ebiten.SetCursorShape(next)
	}
}
