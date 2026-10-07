package reader

import (
	"os"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// rotate switches to a new file at the same path after the old handle is
// drained, so every committed byte was read before the generation changes.
func (r *File) rotate(b *Batch) error {
	if r.f != nil && !r.drained() {
		b.More = true

		return nil
	}

	b.Records = append(b.Records, r.flushRaw()...)

	if r.f != nil {
		_ = r.f.Close()
		r.f = nil
	}

	r.gen++
	r.pos = 0
	r.reset()
	r.drain = -1
	r.identity = ""
	r.opts.HeadHash = 0
	b.Rotated = true

	if err := r.open(); err != nil {
		if os.IsNotExist(err) {
			r.missing, r.state = true, entry.StateMissing
			b.Missing = true

			return nil
		}

		return err
	}

	return nil
}

// drained reports whether the open handle was read up to the size it had
// when the path was replaced.
func (r *File) drained() bool {
	if r.f == nil {
		return true
	}

	if r.drain < 0 {
		if s, err := r.f.Stat(); err == nil {
			r.drain = s.Size()
		}
	}

	return r.pos >= r.drain
}
