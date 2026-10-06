package ui

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/collector"
)

func TestCollectorModeSwitch(t *testing.T) {
	ctx := context.Background()
	mgr := dummyManager(t)

	app, err := New(ctx, Config{Source: NewCollector(mgr), noPump: true, Width: 1024, Height: 700})
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() { _ = app.Close() })

	if err := app.handleClick(ctx, actMode); err != nil {
		t.Fatal(err)
	}

	if mgr.Mode() != collector.ModeLive {
		t.Fatalf("manager mode = %q", mgr.Mode())
	}

	if err := app.handleClick(ctx, actMode); err != nil {
		t.Fatal(err)
	}

	if mgr.Mode() != collector.ModeDummy {
		t.Fatalf("manager mode = %q", mgr.Mode())
	}
}
