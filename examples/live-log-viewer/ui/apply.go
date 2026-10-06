package ui

import (
	"fmt"
	"time"
)

// apply folds one worker result into the view state. It returns true when
// the next frame should render again.
func (a *App) apply(o out, now time.Time) bool {
	switch o.kind {
	case reqPage:
		return a.applyPage(o)
	case reqTail:
		return a.applyTail(o)
	case reqSources:
		return a.applySources(o)
	case reqDetail:
		return a.applyDetail(o)
	case reqSettings:
		return a.applySettings(o)
	case reqSave:
		if o.err != nil {
			a.setNote("settings were not saved: "+o.err.Error(), now)
		}

		return o.err != nil
	case reqExport:
		return a.applyExport(o, now)
	}

	return false
}

// applyPage installs a page result unless a newer request superseded it.
func (a *App) applyPage(o out) bool {
	if o.gen != a.pageGen.Load() {
		return false
	}

	if o.err != nil {
		a.setNote("page load failed: "+o.err.Error(), time.Now())

		return true
	}

	a.pager.Load(o.page)
	if n := a.pager.Len(); n > 0 {
		a.follow.Seen(a.pager.Entries[n-1].ID)
		a.tailAfter.Store(a.follow.LastSeen)
	}

	if o.page.Skipped > 0 {
		a.setNote(fmt.Sprintf("%d older rows were removed by retention", o.page.Skipped), time.Now())
	}

	a.land()
	return true
}

// land puts the viewport where the load intent wants it.
func (a *App) land() {
	switch a.intent {
	case intentNewest:
		a.follow.SetFollow(true)
		a.follow.ClearUnread()
		a.scrollBottom()
	case intentFilter:
		a.scrollBottom()
	case intentOlder:
		a.follow.SetFollow(false)
		a.scrollBottom()
	case intentNewer, intentOldest:
		a.follow.SetFollow(false)
		a.scrollTop()
	case intentAnchor:
		a.scrollTo(a.pager.AnchorOffset(HeaderH))
	}
}
