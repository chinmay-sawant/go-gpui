package ui

import (
	"context"
	"time"
)

// reqKind tells the worker what to run.
type reqKind uint8

const (
	reqPage reqKind = iota
	reqTail
	reqDetail
	reqExport
	reqSources
	reqSettings
	reqSave
)

// req is one worker request. The UI loop never waits for it.
type req struct {
	kind     reqKind
	gen      uint64
	ctx      context.Context
	q        Query
	id       int64
	path     string
	after    int64
	limit    int
	settings Settings
}

// out is one worker result. Snapshots carry the generation they were
// computed for; the tick discards a stale one.
type out struct {
	kind     reqKind
	gen      uint64
	id       int64
	after    int64
	page     PageResult
	entries  []Entry
	total    int
	detail   Detail
	sources  []SourceInfo
	count    int
	path     string
	settings Settings
	err      error
}

// run is the worker loop: one serialized request at a time plus the tail
// poll. A slow query delays later snapshots but never the UI loop.
func (a *App) run() {
	defer a.wg.Done()

	ticker := time.NewTicker(a.pollEvery)
	defer ticker.Stop()

	for {
		select {
		case <-a.ctx.Done():
			return
		case r := <-a.reqs:
			a.serve(r)
		case <-ticker.C:
			a.poll()
		}
	}
}
