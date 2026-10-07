package ui

import (
	"image"

	"github.com/chinmay-sawant/ownframe"
)

// barColor is the progress fill, #2f7fe0, in both themes.
var barColor = [3]float64{47.0 / 255, 127.0 / 255, 224.0 / 255}

// bindings holds the retained operations of one display generation. The row
// slices are index-aligned with view.Active; a nil entry means the row has
// no such operation.
type bindings struct {
	fill   []*ownframe.DisplayOp
	trackW []float64
	pct    []*ownframe.DisplayOp
	speed  []*ownframe.DisplayOp
	eta    []*ownframe.DisplayOp
	sum    [4]*ownframe.DisplayOp
	detail [3]*ownframe.DisplayOp
	foot   *ownframe.DisplayOp
	row    []image.Rectangle
}
