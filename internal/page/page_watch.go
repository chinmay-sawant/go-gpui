package page

import (
	"crypto/sha256"
	"os"
)

type watchKind int

const (
	watchHTML watchKind = iota
	watchTheme
)

// watchPath is one watched file: where it is, the last stat and bytes this
// page accepted, and the stat error already reported.
type watchPath struct {
	path    string
	info    os.FileInfo
	data    []byte
	sum     [sha256.Size]byte
	statErr string
}

// accept records bytes and the stat that read them.
func (w *watchPath) accept(data []byte, info os.FileInfo) {
	w.data = data
	w.sum = sha256.Sum256(data)
	w.info = info
}

// pendingEdit is a read that did not parse, kept for the next poll.
type pendingEdit struct {
	kind watchKind
	data []byte
	info os.FileInfo
}

// watchState is the file watch behind PollReload.
type watchState struct {
	enabled bool
	html    watchPath
	theme   watchPath
	pending *pendingEdit
}

// watchFor builds the watch from the config and the bytes New read.
func watchFor(cfg Config, html []byte, htmlInfo os.FileInfo, theme []byte, themeInfo os.FileInfo) *watchState {
	if cfg.File == "" && cfg.ThemeFile == "" {
		return nil
	}

	w := &watchState{enabled: !cfg.DisableHotReload && watchHosts}
	if cfg.File != "" {
		w.html.path = cfg.File
		w.html.accept(html, htmlInfo)
	}

	if cfg.ThemeFile != "" {
		w.theme.path = cfg.ThemeFile
		w.theme.accept(theme, themeInfo)
	}

	return w
}

func (w *watchState) at(kind watchKind) *watchPath {
	if kind == watchTheme {
		return &w.theme
	}

	return &w.html
}

func (w *watchState) on() bool {
	if w == nil || !w.enabled {
		return false
	}

	return w.html.path != "" || w.theme.path != ""
}
