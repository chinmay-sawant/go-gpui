package reader

import (
	"fmt"
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func dummyMessage(key string, seq int64, r uint64) string {
	var msg string

	switch key {
	case "api":
		msg = fmt.Sprintf("GET /v1/users/%d 200 %dms", r%1000, r%40+1)
	case "worker":
		msg = fmt.Sprintf("job %d finished in %dms", seq/3, r%900+1)
	case "db":
		msg = fmt.Sprintf("checkpoint %d pages wal=%dK", seq/10, r%64)
	default:
		msg = fmt.Sprintf("%s record %d", key, seq)
	}

	switch {
	case seq%71 < 2:
		msg = "connection refused; retrying"
	case seq%73 == 0:
		msg = "panic: runtime error: index out of range"
	case seq%67 == 0:
		msg = "stack for request\n\tat net/http.serverHandler\n\tat main.serve"
	}

	if seq%89 == 0 {
		msg += " — очередь ✓ 日本語 café 🚀"
	}

	if seq%83 == 0 {
		msg += " " + strings.Repeat("payload ", 400)
	}

	if seq%101 == 0 {
		msg += "\x80\x80"
	}

	return msg
}

func severityFor(r uint64) entry.Severity {
	switch n := r % 100; {
	case n < 2:
		return entry.Trace
	case n < 8:
		return entry.Debug
	case n < 70:
		return entry.Info
	case n < 88:
		return entry.Warn
	case n < 97:
		return entry.Error
	default:
		return entry.Fatal
	}
}
