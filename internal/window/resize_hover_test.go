package window

import (
	"context"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// moveScreen keeps one control against the right edge, so a window resize
// moves it out from under or back under a fixed cursor.
type moveScreen struct {
	fakeScreen
	hover   string
	pressed string
}

func (m *moveScreen) Redraw(context.Context) error {
	m.redraws++
	m.boxes = []layout.Box{{ID: "close", X: float64(m.width - 80), Y: 20, W: 60, H: 40}}

	return nil
}

func (m *moveScreen) Hover(_ context.Context, x, y float64) error {
	return m.set(&m.hover, m.hit(x, y))
}

func (m *moveScreen) Press(_ context.Context, x, y float64) error {
	return m.set(&m.pressed, m.hit(x, y))
}

func (m *moveScreen) set(id *string, next string) error {
	if *id == next {
		return nil
	}

	*id = next
	m.redraws++

	return nil
}

func (m *moveScreen) hit(x, y float64) string {
	for _, b := range m.boxes {
		if x >= b.X && x <= b.X+b.W && y >= b.Y && y <= b.Y+b.H {
			return b.ID
		}
	}

	return ""
}
