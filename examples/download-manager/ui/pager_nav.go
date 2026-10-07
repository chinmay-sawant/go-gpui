package ui

// Next moves to the following page. It reports whether a request will
// follow.
func (p *Pager) Next() bool {
	if p.pos+1 < len(p.cursors) {
		p.pos++

		return true
	}

	if p.next == "" {
		return false
	}

	p.cursors = append(p.cursors, p.next)
	p.pos++

	return true
}

// Prev moves to the previous page.
func (p *Pager) Prev() bool {
	if p.pos == 0 {
		return false
	}

	p.pos--

	return true
}

// Rows returns the rows of the accepted page.
func (p *Pager) Rows() []Row { return p.rows }

// Total returns the filtered count.
func (p *Pager) Total() int { return p.total }

// HasNext reports whether a following page exists.
func (p *Pager) HasNext() bool { return p.pos+1 < len(p.cursors) || p.next != "" }

// HasPrev reports whether an earlier page exists.
func (p *Pager) HasPrev() bool { return p.pos > 0 }

// PageNumber returns the 1-based page number.
func (p *Pager) PageNumber() int { return p.pos + 1 }

// Pages returns the known page count, never below the pages visited.
func (p *Pager) Pages() int {
	if p.total <= 0 {
		return p.pos + 1
	}

	n := (p.total + p.limit - 1) / p.limit
	if n < p.pos+1 {
		n = p.pos + 1
	}

	return n
}

// Select stores the selected job ID.
func (p *Pager) Select(id string) { p.sel = id }

// Selected returns the selected job ID.
func (p *Pager) Selected() string { return p.sel }

// Loading reports whether a page answer is still expected.
func (p *Pager) Loading() bool { return p.loading }
