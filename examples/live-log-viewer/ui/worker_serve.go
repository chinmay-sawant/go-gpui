package ui

import "context"

// serve runs one request and sends its result, bounded by the app context.
func (a *App) serve(r req) {
	if r.ctx == nil {
		r.ctx = a.ctx
	}

	ctx, cancel := context.WithTimeout(r.ctx, queryTimeout)
	defer cancel()

	var o out
	o.kind, o.gen, o.id, o.after = r.kind, r.gen, r.id, r.after

	switch r.kind {
	case reqPage:
		o.page, o.err = a.feed.Page(ctx, r.q)
	case reqTail:
		o.entries, o.total, o.err = a.feed.Tail(ctx, r.after, r.limit)
	case reqDetail:
		o.detail, o.err = a.feed.Detail(ctx, r.id)
	case reqExport:
		o.count, o.err = a.feed.Export(ctx, r.q, r.path)
		o.path = r.path
	case reqSources:
		o.sources, o.err = a.feed.Sources(ctx)
	case reqSettings:
		o.settings, o.err = a.feed.Settings(ctx)
	case reqSave:
		o.err = a.feed.SaveSettings(ctx, r.settings)
	}

	select {
	case a.outs <- o:
	case <-a.ctx.Done():
	}
}
