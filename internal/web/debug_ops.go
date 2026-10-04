package web

import "github.com/chinmay-sawant/gowkhtmltopdf/layout"

// debugOpCounts counts the operations that paint, keyed by kind name.
func debugOpCounts(display *layout.Display) map[string]int {
	counts := map[string]int{}

	for i := range display.Ops {
		name := debugOpName(display.Ops[i].Kind)
		if name == "" {
			continue
		}

		counts[name]++
		counts["total"]++
	}

	return counts
}

// debugOpName names one operation kind. Noop, unknown, and link operations
// paint nothing, so they are left out.
func debugOpName(kind layout.DisplayKind) string {
	switch kind {
	case layout.DisplayOpFillRect:
		return "fill"
	case layout.DisplayOpStrokeRect:
		return "stroke"
	case layout.DisplayOpLine:
		return "line"
	case layout.DisplayOpText:
		return "text"
	case layout.DisplayOpBullet:
		return "bullet"
	case layout.DisplayOpImage:
		return "image"
	case layout.DisplayOpGridRun:
		return "grid"
	}

	return ""
}
