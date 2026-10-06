package window

import (
	"time"

	"github.com/chinmay-sawant/ownframe/internal/host"
)

// The hooks below point at the newest shell via wirePerf. Each hook reports
// !ok until a Perf-enabled shell wires them, so the panel shows "-" by
// default and live numbers once a developer opts in.
var (
	devFramePerf    func() (avg, p95, p99 time.Duration, long uint64, ok bool)
	devPipelinePerf func(st host.Stats) (tpl, lay, dsp, pnt time.Duration, dOps, dReg, chOps int, ok bool)
	devRuntimePerf  func() (alloc, heap, rss uint64, gr int, ok bool)
)

func addPerfRows(rs *[]devRow, w float64, h func(string), n func(string, any)) {
	h("Performance")
	if devFramePerf == nil {
		for _, l := range []string{"Frame avg", "Frame p95", "Frame p99", "Long frames"} {
			n(l, "-")
		}
		return
	}
	avg, p95, p99, long, ok := devFramePerf()
	if !ok {
		for _, l := range []string{"Frame avg", "Frame p95", "Frame p99", "Long frames"} {
			n(l, "-")
		}
		return
	}
	n("Frame avg", devMS(avg))
	n("Frame p95", devMS(p95))
	n("Frame p99", devMS(p99))
	n("Long frames", long)
}

func addPipelinePerfRows(rs *[]devRow, w float64, h func(string), n func(string, any), st host.Stats) {
	h("Pipeline detail")
	if devPipelinePerf == nil {
		for _, l := range []string{"Template", "Layout", "Display list", "Paint", "Dirty ops", "Dirty regions", "Changed ops"} {
			n(l, "-")
		}
		return
	}
	tpl, lay, dsp, pnt, dOps, dReg, chOps, ok := devPipelinePerf(st)
	if !ok {
		for _, l := range []string{"Template", "Layout", "Display list", "Paint", "Dirty ops", "Dirty regions", "Changed ops"} {
			n(l, "-")
		}
		return
	}
	n("Template", devMS(tpl))
	n("Layout", devMS(lay))
	n("Display list", devMS(dsp))
	n("Paint", devMS(pnt))
	n("Dirty ops", dOps)
	n("Dirty regions", dReg)
	n("Changed ops", chOps)
}
