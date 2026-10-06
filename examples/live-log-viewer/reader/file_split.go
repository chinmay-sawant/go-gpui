package reader

import (
	"bytes"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

// consume splits a chunk into records, stopping at max records so one large
// chunk cannot blow past the batch limit. It returns how many bytes it used;
// the caller rewinds the rest of the chunk.
func (r *File) consume(data []byte, max int) ([]entry.RawRecord, int) {
	total := len(data)
	cur := r.pos - int64(total)
	if r.started {
		cur = r.bufStart
	}

	var out []entry.RawRecord

	for len(data) > 0 && len(out) < max {
		i := bytes.IndexByte(data, '\n')
		if i < 0 {
			r.appendBuf(data, cur)

			return out, total
		}

		r.appendBuf(data[:i], cur)
		out = append(out, r.complete(cur, cur+int64(i)+1))
		cur += int64(i) + 1
		data = data[i+1:]
	}

	return out, total - len(data)
}

// appendBuf keeps at most MaxRecord bytes and counts the rest as skipped, so
// one huge line cannot grow memory without bound.
func (r *File) appendBuf(data []byte, start int64) {
	if !r.started {
		r.bufStart, r.started = start, true
	}

	room := r.pol.MaxRecord - len(r.buf)
	if room > len(data) {
		room = len(data)
	}

	if room < 0 {
		room = 0
	}

	if room < len(data) {
		r.cut = true
	}

	r.buf = append(r.buf, data[:room]...)
	r.dropped += len(data) - room
}

func (r *File) complete(start, end int64) entry.RawRecord {
	rec := r.buildRecord(r.buf, start, int(end-start), false)
	r.reset()

	return rec
}

func (r *File) flushRaw() []entry.RawRecord {
	if !r.started && len(r.buf) == 0 {
		return nil
	}

	rec := r.buildRecord(r.buf, r.bufStart, len(r.buf)+r.dropped, true)
	r.reset()

	return []entry.RawRecord{rec}
}
