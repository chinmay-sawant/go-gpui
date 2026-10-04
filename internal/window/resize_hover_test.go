package window

import (
	"context"
	"testing"

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

func TestResizeReResolvesHoverAndPress(t *testing.T) {
	t.Parallel()

	app := &moveScreen{fakeScreen: fakeScreen{
		width: 600, height: 500,
		minW: 1, minH: 1,
		maxW: 2560, maxH: 2560,
	}}
	app.boxes = []layout.Box{{ID: "close", X: 520, Y: 20, W: 60, H: 40}}
	app.hover = "close"
	app.pressed = "close"

	game := NewGame(context.Background(), app)
	shell := game.(*shell)
	shell.cursorX, shell.cursorY = 550, 40
	shell.mouseDown = true

	shell.Layout(300, 500)
	if err := shell.resize(); err != nil {
		t.Fatal(err)
	}

	if app.hover != "" || app.pressed != "" {
		t.Fatalf("hover = %q, pressed = %q after the control moved away", app.hover, app.pressed)
	}

	// Widen back: the control returns under the cursor.
	shell.Layout(600, 500)
	if err := shell.resize(); err != nil { // moving, inside the throttle
		t.Fatal(err)
	}

	if err := shell.resize(); err != nil { // settled
		t.Fatal(err)
	}

	if app.hover != "close" || app.pressed != "close" {
		t.Fatalf("hover = %q, pressed = %q after the control returned", app.hover, app.pressed)
	}
}
