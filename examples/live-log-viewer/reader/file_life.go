package reader

import "github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"

// Flush emits the buffered unterminated record, if any.
func (r *File) Flush() []entry.RawRecord { return r.flushRaw() }

// Replayable reports that a file keeps its data and can lag without loss.
func (r *File) Replayable() bool { return true }

// Close releases the file handle.
func (r *File) Close() error {
	r.closed = true

	if r.f != nil {
		err := r.f.Close()
		r.f = nil

		return err
	}

	return nil
}
