package store

import (
	"strings"
	"testing"
)

func TestFileDSNEncodesUnixPaths(t *testing.T) {
	cases := []struct{ path, prefix string }{
		{"/home/u/tetris.db", "file:///home/u/tetris.db?"},
		{"/home/u/a b#c?.db", "file:///home/u/a%20b%23c%3F.db?"},
		{"/home/u/日本.db", "file:///home/u/%E6%97%A5%E6%9C%AC.db?"},
	}

	for _, tc := range cases {
		got := fileDSN(tc.path)
		if !strings.HasPrefix(got, tc.prefix) {
			t.Fatalf("fileDSN(%q) = %q, want prefix %q", tc.path, got, tc.prefix)
		}
	}
}

func TestFileDSNCarriesTheConnectionPragmas(t *testing.T) {
	dsn := fileDSN("/tmp/x.db")

	for _, want := range []string{"busy_timeout", "foreign_keys", "synchronous"} {
		if !strings.Contains(dsn, want) {
			t.Fatalf("DSN misses %s: %q", want, dsn)
		}
	}

	if strings.Contains(dsn, "journal_mode") {
		t.Fatalf("DSN must not switch the journal before the version check: %q", dsn)
	}
}

func TestMemoryDSNUsesTheMemoryPath(t *testing.T) {
	dsn := memoryDSN()

	if !strings.HasPrefix(dsn, Memory+"?") {
		t.Fatalf("memory DSN %q", dsn)
	}

	if !strings.Contains(dsn, "busy_timeout") {
		t.Fatalf("memory DSN misses the busy timeout: %q", dsn)
	}
}

func TestSlashAbsAddsALeadingSlash(t *testing.T) {
	if got := slashAbs("/a/b"); got != "/a/b" {
		t.Fatalf("slashAbs = %q", got)
	}

	if got := slashAbs("a/b"); got != "/a/b" {
		t.Fatalf("slashAbs = %q", got)
	}
}
