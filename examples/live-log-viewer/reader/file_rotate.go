package reader

import (
	"os"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// checkPath notices truncation and replacement before new bytes are read.
// A replaced file is drained to its old size first so no committed bytes
// are skipped.
func (r *File) checkPath(b *Batch) error {
	fi, err := os.Stat(r.opts.Path)
	if err != nil {
		if os.IsNotExist(err) {
			r.missing, r.state = true, entry.StateMissing
			b.Missing = true

			return nil
		}

		return err
	}

	r.missing = false
	id, idErr := pathIdentity(r.opts.Path)

	if r.identity == "" && idErr == nil {
		r.identity = id
	}

	if idErr == nil && id != "" && r.identity != "" && id != r.identity {
		return r.rotate(b)
	}

	if fi.Size() < r.pos {
		b.Records = append(b.Records, r.flushRaw()...)

		r.gen++
		r.pos = 0
		r.reset()
		r.opts.HeadHash = 0
		b.Rotated = true
		r.refreshHead()
	}

	r.size = fi.Size()

	if r.f == nil {
		if err := r.open(); err != nil {
			if os.IsNotExist(err) {
				r.missing, r.state = true, entry.StateMissing
				b.Missing = true

				return nil
			}

			return err
		}
	}

	return nil
}
