package window

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// devFrameRows returns the Frame tab: uppercase headers and right-aligned rows.
func (s *shell) devFrameRows(width float64) []devRow {
	st := s.dev.stats
	appW, appH := s.app.Size()
	rs := []devRow{}
	h := func(s string) { rs = append(rs, devRow{}, devRow{line: devText(strings.ToUpper(s), devDim)}) }
	n := func(l string, v any) { rs = append(rs, devFrameRow(width, l, fmt.Sprint(v), devNumInk)) }
	b := func(l string, v bool) { rs = append(rs, devFrameRow(width, l, fmt.Sprint(v), devBoolInk)) }

	h("Window")
	n("Window", fmt.Sprintf("%dx%d", s.screenW, s.screenH))
	n("Page", fmt.Sprintf("%dx%d", appW, appH))
	n("Scroll", fmt.Sprintf("%d,%d", s.scrollX, s.scrollY))
	b("Stretched", s.stretched())
	b("Fallback", s.fallback)
	n("Sequence", s.seq)
	h("Rendering")
	n("FPS", fmt.Sprintf("%.1f", ebiten.ActualFPS()))
	n("TPS", fmt.Sprintf("%.1f", ebiten.ActualTPS()))
	n("Frame", devMS(s.dev.frame))
	n("Draw", devMS(s.dev.draw))
	h("Pipeline")
	n("Redraws", st.Redraws)
	n("Parses", st.Parses)
	n("Cascades", st.Cascades)
	n("Layouts", st.Layouts)
	n("Repaints", st.Repaints)
	n("Relayouts", s.commits)
	n("Skipped", s.skipped)
	n("Boxes", st.Boxes)
	n("Ops", st.Ops)
	n("Last redraw", devMS(st.LastRedraw))
	n("Last draw", devMS(st.LastDraw))
	h("Reload")
	n("Reloads", st.Reloads)
	ink := devErrInk
	if st.LastReloadError == "" {
		ink = devDim
	}
	rs = append(rs, devFrameRow(width, "Reload error", devErr(st.LastReloadError), ink))

	return rs[1:]
}

// devFrameRow builds one right-aligned row.
func devFrameRow(w float64, l, v string, ink color.RGBA) devRow {
	lw, _ := text.Measure(l, badgeFace, 0)
	vw, _ := text.Measure(v, badgeFace, 0)
	sw, _ := text.Measure(" ", badgeFace, 0)
	n := max(1, int((w-lw-vw)/sw))
	return devRow{line: devLine{spans: []devSpan{{l, devDim, 0}, {strings.Repeat(" ", n), devPunctInk, 0}, {v, ink, 0}}}}
}
