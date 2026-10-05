package benchutil

import (
	"encoding/json"
	"time"

	"github.com/chinmay-sawant/go-gpui"
)

// Dump encodes one perf snapshot: the app name, a UTC timestamp, and the
// page Stats holding counters, pipeline stages, dirty counts, and allocs.
func Dump(name string, page *gpui.Page) ([]byte, error) {
	doc := map[string]any{
		"name":  name,
		"time":  time.Now().UTC().Format(time.RFC3339Nano),
		"stats": page.Stats(),
	}

	return json.MarshalIndent(doc, "", "  ")
}
