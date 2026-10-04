package page

import (
	"context"
	"crypto/sha256"
	"os"
)

// PollReload stats and rereads the watched files. It reports whether the
// page changed, and returns the first error seen. A read that does not parse
// keeps the last good template or theme and is retried on the next poll.
func (p *Page) PollReload(ctx context.Context) (bool, error) {
	if err := useContext(ctx); err != nil {
		return false, err
	}

	if !p.watch.on() {
		return false, nil
	}

	applied := false
	var first error

	if p.watch.pending != nil {
		applied, first = p.retryPending()
	} else {
		for _, kind := range []watchKind{watchHTML, watchTheme} {
			ok, err := p.pollSource(kind)
			if ok {
				applied = true
			}

			if err != nil && first == nil {
				first = err
			}
		}
	}

	if !applied {
		return false, first
	}

	if err := p.Redraw(ctx); err != nil {
		return false, err
	}

	p.resolveStates()

	return true, first
}

// pollSource rereads one path when its stat changed and its hash differs.
func (p *Page) pollSource(kind watchKind) (bool, error) {
	wp := p.watch.at(kind)
	if wp.path == "" {
		return false, nil
	}

	info, err := os.Stat(wp.path)
	if err != nil {
		return false, p.statNote(wp, err)
	}

	wp.statErr = ""
	if sameStat(info, wp.info) {
		return false, nil
	}

	data, err := os.ReadFile(wp.path)
	if err != nil {
		return false, p.statNote(wp, err)
	}

	if wp.info != nil && wp.sum == sha256.Sum256(data) {
		wp.info = info

		return false, nil
	}

	return p.applySource(kind, data, info)
}

// retryPending retries the bytes a parse error kept. A later stat change
// drops the pending bytes and takes the new read.
func (p *Page) retryPending() (bool, error) {
	pend := p.watch.pending
	wp := p.watch.at(pend.kind)

	info, err := os.Stat(wp.path)
	if err != nil {
		return false, p.statNote(wp, err)
	}

	wp.statErr = ""
	if !sameStat(info, pend.info) {
		p.watch.pending = nil

		return p.pollSource(pend.kind)
	}

	return p.applySource(pend.kind, pend.data, info)
}
