package ui

import "testing"

func TestClickThemePersists(t *testing.T) {
	app, back := newTestApp(t)

	clickAction(t, app, "theme")

	if !app.view.Dark {
		t.Fatal("dark mode not set")
	}

	if !back.dark {
		t.Fatal("theme not persisted")
	}

	clickAction(t, app, "theme")

	if app.view.Dark || back.dark {
		t.Fatal("light mode not restored")
	}
}

func TestClickFilterResetsPager(t *testing.T) {
	app, back := newTestApp(t)

	clickAction(t, app, "filter-failed")

	if app.pager.Filter() != FilterFailed {
		t.Fatalf("filter = %v", app.pager.Filter())
	}

	if len(back.pages) == 0 {
		t.Fatal("no page request after the filter change")
	}
}

func TestCloseStopsBackend(t *testing.T) {
	app, back := newTestApp(t)

	if err := app.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}

	if !back.closed {
		t.Fatal("backend not closed")
	}
}
