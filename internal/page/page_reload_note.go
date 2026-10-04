package page

// noteReload records one successful reload for host.Stats. err is the first
// error seen in the same poll, nil when the poll was clean.
func (p *Page) noteReload(err error) {
	p.stats.reloads++

	if err != nil {
		p.stats.reloadErr = err.Error()

		return
	}

	p.stats.reloadErr = ""
}

// noteReloadError records a failed poll that applied nothing.
func (p *Page) noteReloadError(err error) {
	if err != nil {
		p.stats.reloadErr = err.Error()
	}
}
