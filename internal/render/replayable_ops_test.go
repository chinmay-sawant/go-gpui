package render

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

func TestReplayableAcceptsRoundedStroke(t *testing.T) {
	t.Parallel()

	display := &layout.Display{
		Ops: []layout.DisplayOp{{Kind: layout.DisplayOpStrokeRect, Radius: 4}},
	}

	if !Replayable(display) {
		t.Fatal("rounded stroke should replay")
	}
}

func TestReplayableAcceptsMaskedStroke(t *testing.T) {
	t.Parallel()

	display := &layout.Display{
		Ops: []layout.DisplayOp{
			{Kind: layout.DisplayOpStrokeRect, Radius: 4, StrokeMask: 1},
		},
	}

	if !Replayable(display) {
		t.Fatal("masked stroke should replay")
	}
}

func TestReplayableAcceptsEllipticalStroke(t *testing.T) {
	t.Parallel()

	display := &layout.Display{
		Ops: []layout.DisplayOp{
			{Kind: layout.DisplayOpStrokeRect, Radius: 4, RadiusY: 2},
		},
	}

	if !Replayable(display) {
		t.Fatal("elliptical stroke should replay")
	}
}

func TestReplayableRejectsUnknownStrokeMask(t *testing.T) {
	t.Parallel()

	display := &layout.Display{
		Ops: []layout.DisplayOp{
			{Kind: layout.DisplayOpStrokeRect, Radius: 4, StrokeMask: 16},
		},
	}

	if Replayable(display) {
		t.Fatal("unknown mask should not replay")
	}
}

func TestReplayableRejectsImageWithoutBytes(t *testing.T) {
	t.Parallel()

	display := &layout.Display{
		Ops: []layout.DisplayOp{{Kind: layout.DisplayOpImage}},
	}

	if Replayable(display) {
		t.Fatal("image without bytes should not replay")
	}
}

func TestReplayableAcceptsTransformedImage(t *testing.T) {
	t.Parallel()

	source := `<html><body><div style="width:20px;height:20px;` +
		`background:linear-gradient(red,blue);transform:rotate(20deg)"></div></body></html>`

	display, err := DisplayList(context.Background(), source, 320, 200)
	if err != nil {
		t.Fatal(err)
	}

	for _, op := range display.Ops {
		if op.Kind != layout.DisplayOpImage || !op.XformSet {
			continue
		}

		image := &layout.Display{Ops: []layout.DisplayOp{op}}
		if !Replayable(image) {
			t.Fatal("transformed image should replay")
		}

		return
	}

	t.Fatal("transformed svg produced no transformed image op")
}

func TestReplayableAcceptsImageWithBytes(t *testing.T) {
	t.Parallel()

	source := `<html><body><svg width="10" height="10">` +
		`<rect width="10" height="10" fill="red"/></svg></body></html>`

	display, err := DisplayList(context.Background(), source, 320, 200)
	if err != nil {
		t.Fatal(err)
	}

	for _, op := range display.Ops {
		if op.Kind != layout.DisplayOpImage {
			continue
		}

		image := &layout.Display{Ops: []layout.DisplayOp{op}}
		if !Replayable(image) {
			t.Fatal("image with bytes should replay")
		}

		return
	}

	t.Fatal("inline svg produced no image op")
}

func TestReplayableAcceptsLetterSpacing(t *testing.T) {
	t.Parallel()

	source := `<html><body><p style="letter-spacing:2px">hi</p></body></html>`

	display, err := DisplayList(context.Background(), source, 320, 200)
	if err != nil {
		t.Fatal(err)
	}

	for _, op := range display.Ops {
		if op.Kind != layout.DisplayOpText || op.LetterSpacing == 0 {
			continue
		}

		spaced := &layout.Display{Ops: []layout.DisplayOp{op}}
		if !Replayable(spaced) {
			t.Fatal("letter-spaced text should replay")
		}

		return
	}

	t.Fatal("no letter-spaced text op")
}
