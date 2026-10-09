package window

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// devFrameRow builds one right-aligned row.
func devFrameRow(w float64, l, v string, ink color.RGBA) devRow {
	lw, _ := text.Measure(l, badgeFace, 0)
	vw, _ := text.Measure(v, badgeFace, 0)
	sw, _ := text.Measure(" ", badgeFace, 0)
	n := max(1, int((w-lw-vw)/sw))
	return devRow{line: devLine{spans: []devSpan{{l, devDim, 0}, {strings.Repeat(" ", n), devPunctInk, 0}, {v, ink, 0}}}}
}

// addMemoryRows prints the MEMORY section. A shell without Perf reports
// !ok, so end users see dashes and developers see live runtime numbers.
func (s *shell) addMemoryRows(rs *[]devRow, w float64, h func(string), n func(string, any)) {
	h("Memory")
	if s.perfHooks.devRuntimePerf == nil {
		for _, l := range []string{"Alloc/frame", "Go heap", "RSS", "Goroutines"} {
			n(l, "-")
		}
		return
	}
	alloc, heap, rss, gr, ok := s.perfHooks.devRuntimePerf()
	if !ok {
		for _, l := range []string{"Alloc/frame", "Go heap", "RSS", "Goroutines"} {
			n(l, "-")
		}
		return
	}
	n("Alloc/frame", devBytes(alloc))
	n("Go heap", devBytes(heap))
	n("RSS", devBytes(rss))
	n("Goroutines", gr)
}

// devBytes prints a byte count for the panel.
func devBytes(b uint64) string {
	if b < 1024 {
		return fmt.Sprintf("%dB", b)
	}
	if b < 1024*1024 {
		return fmt.Sprintf("%.1fKB", float64(b)/1024)
	}
	return fmt.Sprintf("%.1fMB", float64(b)/(1024*1024))
}
