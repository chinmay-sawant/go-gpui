package ui

import "sort"

// IndexOf returns the index of id, or the index of the first entry whose ID
// is above id. ok is false only when the page is empty.
func (p *Pager) IndexOf(id int64) (int, bool) {
	n := len(p.Entries)
	if n == 0 {
		return 0, false
	}

	i := sort.Search(n, func(i int) bool { return p.Entries[i].ID >= id })
	if i < n {
		return i, true
	}

	return n - 1, true
}

// AnchorIndex returns the row index the reading anchor resolves to.
func (p *Pager) AnchorIndex() int {
	if len(p.Entries) == 0 {
		return 0
	}

	i, _ := p.IndexOf(p.AnchorID)

	return i
}

// AnchorOffset is the scroll offset that puts the anchor row back at the
// viewport top, clamped to the zero floor.
func (p *Pager) AnchorOffset(listTop int) int {
	if len(p.Entries) == 0 {
		return 0
	}

	off := listTop + p.AnchorIndex()*RowH + p.AnchorDelta
	if off < 0 {
		return 0
	}

	return off
}

// SetAnchorFromOffset records the row at the viewport top so a later
// relayout can put it back. listTop is the y of the first row.
func (p *Pager) SetAnchorFromOffset(offsetY, listTop int) {
	if len(p.Entries) == 0 {
		return
	}

	delta := offsetY - listTop
	idx := 0
	if delta >= 0 {
		idx = delta / RowH
	} else {
		delta = 0
	}

	if idx >= len(p.Entries) {
		idx = len(p.Entries) - 1
		delta = 0
	}

	p.AnchorID = p.Entries[idx].ID
	p.AnchorDelta = delta % RowH
}

// FindAnchor returns the ID to restore, or zero when there is none.
func (p *Pager) FindAnchor() int64 {
	if len(p.Entries) == 0 {
		return 0
	}

	if p.AnchorID == 0 {
		return p.Entries[0].ID
	}

	return p.AnchorID
}
