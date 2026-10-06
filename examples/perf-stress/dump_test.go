// Command perfstress_test dumps headless perf snapshots as JSON. It
// redraws the stress dashboard and each benchutil app once, then writes one
// combined document, so scripts/perf-dump.sh archives a run headless.
package main

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/perf-benchmarks/benchutil"
)

// TestDump redraws the stress app and every bench app headless and writes
// the combined Stats snapshots to $PERF_DUMP_DIR/dump-<ts>.json.
func TestDump(t *testing.T) {
	ctx := context.Background()

	stress, err := newApp(240)
	if err != nil {
		t.Fatal(err)
	}

	normal, err := benchutil.NewNormal()
	if err != nil {
		t.Fatal(err)
	}

	large, err := benchutil.NewLarge(1000)
	if err != nil {
		t.Fatal(err)
	}

	flap, err := benchutil.NewFlap()
	if err != nil {
		t.Fatal(err)
	}

	apps := []struct {
		name   string
		page   *ownframe.Page
		redraw func(context.Context) error
	}{
		{"stress-240", stress.page, stress.page.Redraw},
		{"normal", normal.Page(), normal.Redraw},
		{"large-1000", large.Page(), large.Redraw},
		{"flappy", flap.Page(), flap.Redraw},
	}

	var docs []json.RawMessage

	for _, a := range apps {
		if err := a.redraw(ctx); err != nil {
			t.Fatal(err)
		}

		b, err := benchutil.Dump(a.name, a.page)
		if err != nil {
			t.Fatal(err)
		}

		docs = append(docs, b)
	}

	out, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	dir := os.Getenv("PERF_DUMP_DIR")
	if dir == "" {
		dir = filepath.FromSlash("../../temp/perf-dumps")
	}

	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}

	p := filepath.Join(dir, "dump-"+time.Now().UTC().Format("20060102T150405Z")+".json")
	if err := os.WriteFile(p, append(out, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}

	t.Logf("wrote %s", p)
}
