package window

import "image"

// repaintMode is what one frame does with the persistent buffer.
type repaintMode int

const (
	// repaintFull replays the list straight to the screen. It is the path for
	// a page that cannot report dirty rects.
	repaintFull repaintMode = iota
	// repaintBuffer repaints the whole buffer.
	repaintBuffer
	// repaintRect repaints one dirty rect in the buffer.
	repaintRect
	// repaintBlit leaves the buffer alone. Its pixels are current.
	repaintBlit
)

// repaintPlan is the buffer work for one frame.
type repaintPlan struct {
	mode repaintMode
	rect image.Rectangle
}

// planRepaint decides the frame's buffer work for content contentW x contentH.
// The content can be taller or wider than the laid-out canvas when a fixed
// root's children overflow, so the buffer and the full rect use the content,
// not the canvas. A dirty rect that reaches outside the content is clamped;
// one that clamps away, or an empty or false rect, forces the safe full
// rebuild instead of a silent skip.
func planRepaint(contentW, contentH int, st partialState, gen uint64, rect image.Rectangle, ok bool) repaintPlan {
	full := image.Rect(0, 0, contentW, contentH)
	if full.Empty() || st.buf == nil || st.width != contentW || st.height != contentH {
		return repaintPlan{mode: repaintBuffer, rect: full}
	}

	if !ok {
		if st.gen == gen {
			return repaintPlan{mode: repaintBlit, rect: full}
		}

		return repaintPlan{mode: repaintBuffer, rect: full}
	}

	clamped := rect.Intersect(full)
	if clamped.Empty() || clamped == full {
		return repaintPlan{mode: repaintBuffer, rect: full}
	}

	return repaintPlan{mode: repaintRect, rect: clamped}
}
