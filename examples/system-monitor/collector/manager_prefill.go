package collector

// prefillLocked loads the dummy fixture history into the graph buffers.
func (m *Manager) prefillLocked() {
	d, ok := m.source.(*Dummy)
	if !ok {
		return
	}

	hist := d.Historical(m.start, m.opts.History, m.opts.HistoryStep)
	if len(hist) == 0 {
		return
	}

	for _, s := range hist {
		m.pushRingsLocked(s)
	}

	last := hist[len(hist)-1]
	m.reading = last
	m.haveRead = true
	m.prev = last
}
