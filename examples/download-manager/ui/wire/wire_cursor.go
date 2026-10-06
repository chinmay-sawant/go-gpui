package wire

import (
	"strings"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/store"
)

// encodeCursor writes the opaque keyset cursor: RFC 3339 time, a bar, and
// the stable job ID.
func encodeCursor(c store.Cursor) string {
	if c.ID == "" && c.UpdatedAt.IsZero() {
		return ""
	}

	return c.UpdatedAt.UTC().Format(time.RFC3339Nano) + "|" + c.ID
}

// decodeCursor parses one cursor. A malformed value starts at the newest
// row instead of failing the request.
func decodeCursor(s string) store.Cursor {
	if s == "" {
		return store.Cursor{}
	}

	at, id, ok := strings.Cut(s, "|")
	if !ok {
		return store.Cursor{}
	}

	t, err := time.Parse(time.RFC3339Nano, at)
	if err != nil {
		return store.Cursor{}
	}

	return store.Cursor{UpdatedAt: t, ID: id}
}
