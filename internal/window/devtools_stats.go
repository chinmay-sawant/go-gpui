package window

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// devFrameTime records the wall time since the previous Draw call.
func (s *shell) devFrameTime() {
	now := time.Now()
	if !s.dev.lastAt.IsZero() {
		s.dev.frame = now.Sub(s.dev.lastAt)
	}

	s.dev.lastAt = now
}

// devLines formats the panel from the frame numbers and the page stats. The
// panel measures these strings to size itself.
func (s *shell) devLines() []string {
	st := s.dev.stats
	appW, appH := s.app.Size()

	lines := []string{
		"devtools  f12 or ctrl+shift+i",
		s.devOpsLine(),
		fmt.Sprintf("fps %.1f  tps %.1f", ebiten.ActualFPS(), ebiten.ActualTPS()),
		fmt.Sprintf("frame %s  draw %s", devMS(s.dev.frame), devMS(s.dev.draw)),
		fmt.Sprintf("window %dx%d  app %dx%d", s.screenW, s.screenH, appW, appH),
		fmt.Sprintf("scroll %d,%d  stretched %v  fallback %v  seq %d",
			s.scrollX, s.scrollY, s.stretched(), s.fallback, s.seq),
		fmt.Sprintf("redraws %d  parses %d  cascades %d  layouts %d  repaints %d",
			st.Redraws, st.Parses, st.Cascades, st.Layouts, st.Repaints),
		fmt.Sprintf("relayouts %d  skipped %d", s.commits, s.skipped),
		fmt.Sprintf("boxes %d  ops %d  last redraw %s  last draw %s",
			st.Boxes, st.Ops, devMS(st.LastRedraw), devMS(st.LastDraw)),
		fmt.Sprintf("reloads %d  reload error %s", st.Reloads, devErr(st.LastReloadError)),
	}

	if s.dev.haveHov {
		lines = append(lines, "hover  "+s.devBoxDetail(s.dev.hovered))
	}

	if s.dev.havePin {
		lines = append(lines, "pin  "+s.devBoxDetail(s.dev.pinned))
	}

	return lines
}

// devMS formats a duration for the panel. A zero duration reads as zero
// instead of a negative or rounded value.
func devMS(d time.Duration) string {
	if d <= 0 {
		return "0.0ms"
	}

	return fmt.Sprintf("%.1fms", float64(d)/float64(time.Millisecond))
}

func devErr(err string) string {
	if err == "" {
		return "-"
	}

	return err
}
