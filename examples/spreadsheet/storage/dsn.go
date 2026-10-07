package storage

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// Memory is an in-memory database that lasts for the process. Tests use it
// and need no directory.
const Memory = ":memory:"

// File is the database file name inside the data directory.
const File = "spreadsheet.db"

// DefaultDir is os.UserConfigDir()/ownframe/spreadsheet.
func DefaultDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(base, "ownframe", "spreadsheet"), nil
}

// dsn builds the modernc.org/sqlite connection string. A plain path handles
// Windows drive letters, spaces, and Unicode; a path holding ? or # becomes
// a file URI with escaping. Ready pragmas ride on the DSN so every physical
// connection inherits them.
func dsn(path string, memory bool) string {
	params := "_pragma=foreign_keys(ON)&_pragma=busy_timeout(5000)"
	if !memory {
		params += "&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)"
	}

	if memory {
		return path + "?" + params
	}

	if strings.ContainsAny(path, "?#") {
		u := url.URL{
			Scheme:   "file",
			Path:     filepath.ToSlash(path),
			RawQuery: params,
		}

		return u.String()
	}

	return path + "?" + params
}
