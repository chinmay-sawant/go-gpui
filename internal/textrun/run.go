package textrun

import (
	"math"
	"unicode/utf8"

	"github.com/hajimehoshi/ebiten/v2/text/v2"

	"github.com/chinmay-sawant/blinkless/layout"
)

// offsetInLine returns the rune offset nearest xPx among the line's runs.
// xPx is a CSS pixel distance from the display's left edge. ok is false when
// a run's font cannot be measured.
func offsetInLine(ln line, xPx, pxPerPt float64) (int, bool) {
	best, bestBase := 0, 0
	bestDist := math.MaxFloat64
	base := 0
	for i, r := range ln.runs {
		if faceFor(r.op, pxPerPt) == nil {
			return 0, false
		}

		left := r.op.X * pxPerPt
		right := left + runWidth(r.op, pxPerPt)
		if xPx >= left && xPx <= right {
			return base + offsetIn(r.op, xPx-left, pxPerPt), true
		}

		dist := left - xPx
		if xPx > right {
			dist = xPx - right
		}
		if dist < bestDist {
			best, bestBase, bestDist = i, base, dist
		}

		base += utf8.RuneCountInString(drawn(r.op))
	}

	op := ln.runs[best].op

	return bestBase + offsetIn(op, xPx-op.X*pxPerPt, pxPerPt), true
}

// runWidth measures a run in CSS pixels.
func runWidth(op *layout.DisplayOp, pxPerPt float64) float64 {
	face := faceFor(op, pxPerPt)
	if face == nil {
		return op.W * pxPerPt
	}

	w, _ := text.Measure(drawn(op), face, 0)

	return w
}

// offsetIn returns the rune offset nearest dxPx from the run's start.
// dxPx is in CSS pixels.
func offsetIn(op *layout.DisplayOp, dxPx, pxPerPt float64) int {
	runes := []rune(drawn(op))
	if len(runes) == 0 {
		return 0
	}

	face := faceFor(op, pxPerPt)
	if face == nil {
		return len(runes)
	}

	spacing := op.LetterSpacing * pxPerPt
	best, bestDist := 0, math.Abs(dxPx)
	width := 0.0

	for i := 1; i <= len(runes); i++ {
		if spacing != 0 {
			glyph, _ := text.Measure(string(runes[i-1]), face, 0)
			width += glyph + spacing
		} else {
			width, _ = text.Measure(string(runes[:i]), face, 0)
		}

		if d := math.Abs(width - dxPx); d < bestDist {
			best, bestDist = i, d
		}
	}

	return best
}
