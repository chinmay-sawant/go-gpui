package window

import (
	"context"
	"slices"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/ownframe/internal/host"
)

// cursorScreen is a fakeScreen with a scripted hover shape.
type cursorScreen struct {
	*fakeScreen
	shape host.Shape
}

func (f *cursorScreen) CursorShape() host.Shape { return f.shape }

func TestCursorShapeSequence(t *testing.T) {
	t.Parallel()

	app := &cursorScreen{fakeScreen: &fakeScreen{}}
	var calls []ebiten.CursorShapeType

	s := &shell{app: app, ctx: context.Background()}
	s.setCursor = func(shape ebiten.CursorShapeType) {
		calls = append(calls, shape)
	}

	for _, shape := range []host.Shape{host.ShapeText, host.ShapePointer, host.ShapeDefault} {
		app.shape = shape
		s.applyCursor()
	}

	want := []ebiten.CursorShapeType{
		ebiten.CursorShapeText,
		ebiten.CursorShapePointer,
		ebiten.CursorShapeDefault,
	}
	if !slices.Equal(calls, want) {
		t.Fatalf("calls = %v, want %v", calls, want)
	}

	s.applyCursor()

	if len(calls) != 3 {
		t.Fatalf("repeat calls = %d", len(calls))
	}

	if s.cursor != ebiten.CursorShapeDefault {
		t.Fatalf("cursor = %v", s.cursor)
	}
}

func TestCursorMapping(t *testing.T) {
	t.Parallel()

	cases := []struct {
		in   host.Shape
		want ebiten.CursorShapeType
	}{
		{host.ShapeDefault, ebiten.CursorShapeDefault},
		{host.ShapeText, ebiten.CursorShapeText},
		{host.ShapePointer, ebiten.CursorShapePointer},
		{host.ShapeResizeEW, ebiten.CursorShapeEWResize},
		{host.ShapeResizeNS, ebiten.CursorShapeNSResize},
	}

	for _, c := range cases {
		if got := cursorShape(c.in); got != c.want {
			t.Fatalf("cursorShape(%v) = %v", c.in, got)
		}
	}
}

func TestCursorPlainScreen(t *testing.T) {
	t.Parallel()

	s := &shell{app: &fakeScreen{}, ctx: context.Background()}
	s.applyCursor()

	if s.cursor != ebiten.CursorShapeDefault {
		t.Fatalf("cursor = %v", s.cursor)
	}
}
