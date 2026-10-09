package window

import (
	"testing"
)

func TestPerfIsOwnedByItsWindow(t *testing.T) {
	a := newDevShell(newDevScreen())
	a.perf = true
	a.wirePerf()
	before := devFrameText(a.devFrameRows(320))
	b := newDevShell(newDevScreen())
	b.wirePerf()
	if got := devFrameText(a.devFrameRows(320)); got != before {
		t.Fatal("another window replaced this window's performance hooks")
	}
}
