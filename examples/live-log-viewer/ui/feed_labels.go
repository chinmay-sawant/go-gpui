package ui

import (
	"strconv"
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// setLabel caches a source display name for entry rows.
func (f *StoreFeed) setLabel(id entry.SourceID, name string) {
	f.mu.Lock()
	f.labels[id] = name
	f.mu.Unlock()
}

// sourceName returns the cached name for a source ID.
func (f *StoreFeed) sourceName(id entry.SourceID) string {
	f.mu.Lock()
	defer f.mu.Unlock()

	if name, ok := f.labels[id]; ok {
		return name
	}

	return "src-" + strconv.FormatInt(int64(id), 10)
}

// toEntry maps one core entry into the display form. The message keeps its
// newlines; the list slices the first line and the detail pane shows all.
func (f *StoreFeed) toEntry(e entry.Entry) Entry {
	return Entry{
		ID:        int64(e.ID),
		Time:      e.Time,
		TimeRaw:   e.TimeRaw,
		TimeOK:    e.TimeOK,
		Source:    f.sourceName(e.Source),
		SourceKey: strconv.FormatInt(int64(e.Source), 10),
		Severity:  strings.ToLower(e.Severity.String()),
		Text:      e.Message,
		Lines:     lineCount(e.Message),
		Bytes:     e.Bytes,
		Truncated: e.Truncated,
		Partial:   e.Partial,
	}
}
