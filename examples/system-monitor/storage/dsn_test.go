package storage

import (
	"net/url"
	"testing"
)

// TestFileDSN checks the URI shapes for spaces, Unicode, URI punctuation,
// Windows drive letters, and UNC paths.
func TestFileDSN(t *testing.T) {
	posix := fileDSN("/tmp/a b/дб#?.db", false)

	u, err := url.Parse(posix)
	if err != nil {
		t.Fatal(err)
	}
	if u.Scheme != "file" || u.Path != "/tmp/a b/дб#?.db" || u.RawQuery != "" {
		t.Fatalf("posix dsn = %q parsed %+v", posix, u)
	}

	drive := fileURI(splitPath("C:/Users/a b/дб.db", true))
	u, err = url.Parse(drive)
	if err != nil {
		t.Fatal(err)
	}
	if u.Path != "/C:/Users/a b/дб.db" || u.Host != "" {
		t.Fatalf("drive dsn = %q parsed %+v", drive, u)
	}

	host, path := splitPath("//server/share/a b.db", true)
	if host != "server" || path != "/share/a b.db" {
		t.Fatalf("unc = %q %q", host, path)
	}

	if got := fileDSN(":memory:", false); got != ":memory:" {
		t.Fatalf("memory = %q", got)
	}
	if got := fileDSN("file:custom.db?x=1", false); got != "file:custom.db?x=1" {
		t.Fatalf("existing uri = %q", got)
	}
}
