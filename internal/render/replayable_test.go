package render

import (
	"testing"

	"github.com/chinmay-sawant/blinkless/layout"
)

func TestReplayableAcceptsPlainOps(t *testing.T) {
	t.Parallel()

	display := &layout.Display{
		Ops: []layout.DisplayOp{
			{Kind: layout.DisplayOpFillRect, Radius: 3},
			{Kind: layout.DisplayOpLine},
			{Kind: layout.DisplayOpNoop},
			{Kind: layout.DisplayOpLinkURI},
		},
	}

	if !Replayable(display) {
		t.Fatal("plain ops should replay")
	}
}

func TestReplayableRejectsUnsupportedOps(t *testing.T) {
	t.Parallel()

	display := &layout.Display{
		Ops: []layout.DisplayOp{{Kind: layout.DisplayOpStrokeRect, StrokeMask: 16}},
	}

	if Replayable(display) {
		t.Fatal("unknown stroke mask should not replay")
	}

	display.Ops[0].Kind = layout.DisplayOpGridRun

	if Replayable(display) {
		t.Fatal("grid run without segments should not replay")
	}
}

func TestReplayableRequiresACircularFill(t *testing.T) {
	t.Parallel()

	display := &layout.Display{
		Ops: []layout.DisplayOp{{Kind: layout.DisplayOpFillRect, Radius: 4, RadiusY: 2}},
	}

	if Replayable(display) {
		t.Fatal("elliptical fill should not replay")
	}
}

func TestReplayableRequiresAFaceForText(t *testing.T) {
	t.Parallel()

	display := &layout.Display{
		Ops: []layout.DisplayOp{{Kind: layout.DisplayOpText, Text: "hi"}},
	}

	if Replayable(display) {
		t.Fatal("text without a face should not replay")
	}
}
