package ui

// Request returns the request for the current page and advances the
// generation, so any earlier answer is now stale.
func (p *Pager) Request() PageRequest {
	p.gen++
	p.loading = true

	return PageRequest{Gen: p.gen, Filter: p.filter, Before: p.cursors[p.pos], Limit: p.limit}
}

// Accept applies a page response. It reports ok when the response belongs
// to the newest request. When a page after the first comes back empty, the
// walk steps back one page and retry reports that the caller should request
// again.
func (p *Pager) Accept(resp PageResponse) (ok, retry bool) {
	if resp.Gen != p.gen {
		return false, false
	}

	if p.pos > 0 && len(resp.Rows) == 0 {
		p.pos--

		return false, true
	}

	p.rows = resp.Rows
	p.total = resp.Total
	p.next = resp.Next
	p.loading = false

	return true, false
}
