package window

import (
	"strings"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/internal/host"
)

// TestDevPerfRuntimeHook checks the runtime hook values render.
func TestDevPerfRuntimeHook(t *testing.T) {
	s := newDevShell(newDevScreen())
	s.perfHooks.devRuntimePerf = func() (uint64, uint64, uint64, int, bool) {
		return 2048, 3 * 1024 * 1024, 512, 9, true
	}

	got := devFrameText(s.devFrameRows(320))
	for _, want := range []string{"2.0KB", "3.0MB", "512B"} {
		if !strings.Contains(got, want) {
			t.Fatalf("frame rows = %q, want %q", got, want)
		}
	}
}

// TestDevPerfHookNotOK checks !ok hooks fall back to dashes.
func TestDevPerfHookNotOK(t *testing.T) {
	s := newDevShell(newDevScreen())
	s.perfHooks.devFramePerf = func() (time.Duration, time.Duration, time.Duration, uint64, bool) {
		return 0, 0, 0, 0, false
	}
	s.perfHooks.devPipelinePerf = func(st host.Stats) (time.Duration, time.Duration, time.Duration, time.Duration, int, int, int, bool) {
		return 0, 0, 0, 0, 0, 0, 0, false
	}
	s.perfHooks.devRuntimePerf = func() (uint64, uint64, uint64, int, bool) { return 0, 0, 0, 0, false }

	_ = s.devFrameRows(320)
}
