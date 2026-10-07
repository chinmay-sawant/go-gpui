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
	settings Settings
}

// out is one worker result. Snapshots carry the generation they were
// computed for; the tick discards a stale one.
type out struct {
	kind     reqKind
	gen      uint64
	id       int64
	page     PageResult
	detail   Detail
	sources  []SourceInfo
	count    int
	path     string
	settings Settings
	err      error
}

// run is the worker loop: one serialized request at a time. The UI loop
// dispatches; the worker never polls on its own.
func (a *App) run() {
	defer a.wg.Done()

	for {
		select {
		case <-a.ctx.Done():
			return
		case r := <-a.reqs:
			a.serve(r)
		}
	}
}

// queryTimeout bounds one store call on the worker side.
const queryTimeout = 5 * time.Second

// poll dispatches one tail poll when its interval is due.
func (a *App) poll(now time.Time) {
	if a.feed == nil || !a.pollOn.Load() || now.Before(a.tailDue) {
		return
	}

	if a.tailAfter.Load() == 0 {
		return // wait for the first page before counting unread
	}

	a.tailDue = now.Add(a.pollEvery)

	q := a.baseQuery()
	q.AfterID = a.tailAfter.Load()
	q.Limit = TailLimit
	a.send(req{kind: reqTail, gen: a.filters.Gen, ctx: a.ctx, q: q})
}
