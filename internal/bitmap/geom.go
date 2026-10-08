package bitmap

import "github.com/chinmay-sawant/blinkless/layout"

// origin is the op's top-left, or its baseline for text, after a baked
// transform. A translation moves the whole rectangle. Width and height stay.
func origin(op *layout.DisplayOp) (float64, float64) {
	if op.XformSet && !op.Transform().IsIdentity() {
		return op.Transform().Apply(op.X, op.Y)
	}

	return op.X, op.Y
}
