package ui

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStoreFeedNoMatches(t *testing.T) {
	a, _, _ := newStoreApp(t)

	pumpUntil(t, a, func() bool { return a.pager.Len() > 0 }, "first page")

	a.filters.Apply("no-such-text-anywhere")
	a.loadPage(intentFilter)
	pumpUntil(t, a, func() bool { return a.pager.Empty() }, "empty result")

	if err := a.draw(context.Background()); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(a.view.Empty, "No entries match") {
		t.Fatalf("empty text = %q", a.view.Empty)
	}
}

func TestStoreFeedDetailAndExport(t *testing.T) {
	a, _, newest := newStoreApp(t)

	pumpUntil(t, a, func() bool { return a.pager.Len() == PageLimit }, "first page")

	a.openDetail(newest)
	pumpUntil(t, a, func() bool { return a.detailOpen }, "detail")

	if a.detail.ID != newest || len(a.detail.Lines) == 0 {
		t.Fatalf("detail = %+v", a.detail)
	}

	first := a.pager.Entries[0].ID
	a.selected = first
	a.selAnchor = first + 4
	a.startExport()
	pumpUntil(t, a, func() bool { return strings.Contains(a.note, "exported") }, "export")

	files, err := os.ReadDir(a.exportDir)
	if err != nil || len(files) != 1 {
		t.Fatalf("exports = %v %v", files, err)
	}

	data, err := os.ReadFile(filepath.Join(a.exportDir, files[0].Name()))
	if err != nil {
		t.Fatal(err)
	}

	if lines := strings.Count(string(data), "\n"); lines != 5 {
		t.Fatalf("exported lines = %d", lines)
	}
}

func TestStoreFeedSettingsPersist(t *testing.T) {
	a, feed, _ := newStoreApp(t)

	pumpUntil(t, a, func() bool { return a.settingsLoaded }, "settings")

	a.toggleTheme()
	pumpUntil(t, a, func() bool {
		s, err := feed.Settings(context.Background())

		return err == nil && s.Dark
	}, "saved theme")

	if !a.dark {
		t.Fatal("theme did not switch")
	}
}
