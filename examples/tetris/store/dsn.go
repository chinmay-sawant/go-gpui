package store

import (
	"database/sql"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

// sqlDB aliases database/sql so the Store struct stays short.
type sqlDB = sql.DB

// fileDSN builds a file: URI for a database path. Percent encoding keeps
// spaces, Unicode, ?, #, and Windows drive letters intact, and the
// _pragma parameters apply to every physical connection.
func fileDSN(path string) string {
	u := url.URL{Scheme: "file", Path: slashAbs(path)}

	return u.String() + "?" + pragmaQuery()
}

// memoryDSN builds a private in-memory DSN with the same pragmas.
func memoryDSN() string { return Memory + "?" + pragmaQuery() }

// pragmaQuery returns the connection pragmas that do not modify a file: a
// finite busy timeout, foreign keys, and full durability. journal_mode is
// activated separately so a newer schema can be rejected untouched.
func pragmaQuery() string {
	q := url.Values{}
	q.Add("_pragma", "busy_timeout("+strconv.Itoa(busyTimeoutMS)+")")
	q.Add("_pragma", "foreign_keys(1)")
	q.Add("_pragma", "synchronous(FULL)")

	return q.Encode()
}

// slashAbs makes a path absolute in slash form, so a Windows drive path
// becomes /C:/dir and the file: URI has three slashes.
func slashAbs(path string) string {
	p := filepath.ToSlash(path)
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}

	return p
}
