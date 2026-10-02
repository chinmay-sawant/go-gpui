package render

import "github.com/chinmay-sawant/gowkhtmltopdf/layout"

// FillRadii resolves a display fill's corner radii in canvas points, in CSS
// order: top-left, top-right, bottom-right, bottom-left. A missing Y radius
// copies the X radius and a zero X radius forces a zero Y radius, matching
// the engine's resolution. ok is false when a corner is elliptical, because
// the vector replay draws circular arcs only.
func FillRadii(op *layout.DisplayOp) (radii [4]float64, ok bool) {
	rx := [4]float64{op.Radius, op.Radius, op.Radius, op.Radius}
	if op.RadiusTopLeft > 0 || op.RadiusTopRight > 0 ||
		op.RadiusBottomRight > 0 || op.RadiusBottomLeft > 0 {
		rx = [4]float64{
			op.RadiusTopLeft, op.RadiusTopRight,
			op.RadiusBottomRight, op.RadiusBottomLeft,
		}
	}

	ry := [4]float64{op.RadiusY, op.RadiusY, op.RadiusY, op.RadiusY}
	if op.RadiusTopLeftY > 0 || op.RadiusTopRightY > 0 ||
		op.RadiusBottomRightY > 0 || op.RadiusBottomLeftY > 0 {
		ry = [4]float64{
			op.RadiusTopLeftY, op.RadiusTopRightY,
			op.RadiusBottomRightY, op.RadiusBottomLeftY,
		}
	}

	for i := range rx {
		if ry[i] <= 0 {
			ry[i] = rx[i]
		}

		if rx[i] <= 0 {
			rx[i], ry[i] = 0, 0
		}

		if rx[i] != ry[i] {
			return [4]float64{}, false
		}
	}

	return rx, true
}
