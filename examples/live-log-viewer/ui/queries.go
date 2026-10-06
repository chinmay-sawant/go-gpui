package ui

import "context"

// pageIntent is why a page is loading; it decides where the viewport lands
// after the result applies.
type pageIntent uint8

const (
	intentNewest pageIntent = iota // jump to the live tail, follow on
	intentFilter                   // newest matches, keep the follow mode
	intentOlder                    // one page toward older entries
	intentNewer                    // one page toward newer entries
	intentOldest                   // the first page of the frozen set
	intentAnchor                   // refresh in place, keep the anchor row
)

// baseQuery copies the active filters into a page request.
func (a *App) baseQuery() Query {
	q := Query{Limit: PageLimit, Text: a.filters.Text}
	q.Severities = append([]string(nil), a.activeSevs...)

	if a.activeSource != "" {
		q.Sources = []string{a.activeSource}
	}

	return q
}

// loadPage dispatches one keyset page request and supersedes the previous
// one: its cancel runs and its generation stops matching.
func (a *App) loadPage(intent pageIntent) {
	q := a.baseQuery()
	q.MaxID = a.pager.HWM

	switch intent {
	case intentNewest, intentFilter:
		q.MaxID = 0
		a.pager.HWM = 0
	case intentOldest:
		a.freeze()
		q.MaxID = a.pager.HWM
		q.Oldest = true
	case intentOlder:
		a.freeze()
		q.MaxID = a.pager.HWM
		q.BeforeID = a.pager.OlderCursor()

		if q.BeforeID == 0 {
			return
		}
	case intentNewer:
		q.MaxID = a.pager.HWM
		q.AfterID = a.pager.NewerCursor()

		if q.AfterID == 0 {
			return
		}
	}

	a.pageGen.Add(1)
	a.intent = intent

	ctx, cancel := context.WithCancel(a.ctx)
	a.setPageCancel(cancel)
	a.send(req{kind: reqPage, gen: a.pageGen.Load(), ctx: ctx, q: q})
}
