package page

import (
	"runtime"

	"github.com/chinmay-sawant/go-gpui/internal/render"
)

// allocStart samples the allocator before a Redraw; allocUsed diffs it.
func allocStart() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return m
}

// allocUsed is the bytes allocated since allocStart sampled m0.
func allocUsed(m0 runtime.MemStats) uint64 {
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	if m1.TotalAlloc < m0.TotalAlloc {
		return 0
	}

	return m1.TotalAlloc - m0.TotalAlloc
}

// renderState is the input, theme, and image state Redraw and PNG share.
func (p *Page) renderState() render.State {
	state := render.State{Hover: p.hover, Active: p.active, Theme: p.theme, Images: p.imageBytes}
	if p.form != nil {
		state.Focus = p.form.focusID
	}

	return state
}
