package page

import "os"

// SetHotReload turns the file watch on or off. It does nothing without a
// watched file, and nothing on wasm and mobile, where a file is read once.
func (p *Page) SetHotReload(on bool) {
	if p.watch == nil {
		return
	}

	p.watch.enabled = on && watchHosts
}

// Watching reports whether a poll can reread a source file.
func (p *Page) Watching() bool {
	return p.watch.on()
}

func sameStat(a, b os.FileInfo) bool {
	if a == nil || b == nil {
		return false
	}

	return a.ModTime().Equal(b.ModTime()) && a.Size() == b.Size()
}
