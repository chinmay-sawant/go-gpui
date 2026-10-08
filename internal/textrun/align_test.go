package textrun

import (
	"testing"

	"github.com/chinmay-sawant/blinkless/layout"
)

func TestAlignSkipsWrapSpace(t *testing.T) {
	t.Parallel()

	lines := []line{
		{baseline: 0, runs: []run{{op: &layout.DisplayOp{Text: "ab cd"}}}},
		{baseline: 10, runs: []run{{op: &layout.DisplayOp{Text: "ef"}}}},
	}
	bases, ok := align(lines, "ab cd ef")
	if !ok || bases[0] != 0 || bases[1] != 6 {
		t.Fatalf("bases %v ok %v", bases, ok)
	}

	if _, ok := align(lines, "xy"); ok {
		t.Fatal("mismatch aligned")
	}
}

func TestAlignRunsInLine(t *testing.T) {
	t.Parallel()

	lines := []line{{
		baseline: 0,
		runs: []run{
			{op: &layout.DisplayOp{Text: "ab"}},
			{op: &layout.DisplayOp{Text: "cd"}},
		},
	}}
	bases, ok := align(lines, "abcd")
	if !ok || bases[0] != 0 {
		t.Fatalf("bases %v ok %v", bases, ok)
	}
}
