package crash_test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/chinmay-sawant/go-gpui"
)

func TestReportWritesAFile(t *testing.T) {
	app := newApp(t, context.Background())

	dir := t.TempDir()
	gpui.SetCrashDir(dir)
	t.Cleanup(func() { gpui.SetCrashDir("") })

	path, err := app.Report("test reason")
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(path, dir) {
		t.Fatalf("path = %q, want under %q", path, dir)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	if info.IsDir() || info.Size() == 0 {
		t.Fatalf("report file = %v, want a non-empty file", info)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	body := string(raw)
	if !strings.Contains(body, "test reason") {
		t.Fatalf("report does not name the reason: %s", body)
	}
}
