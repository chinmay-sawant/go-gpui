package reader

import "github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"

func (r *File) buildRecord(line []byte, start int64, size int, partial bool) entry.RawRecord {
	rec := entry.RawRecord{
		Path: r.opts.Path, Generation: r.gen, Offset: start,
		Bytes: size, Partial: partial, Skipped: r.dropped,
	}

	body := line
	if n := len(body); n > 0 && body[n-1] == '\r' {
		body = body[:n-1]
	}

	if len(body) > r.pol.MaxRecord {
		rec.Skipped += len(body) - r.pol.MaxRecord
		body = body[:r.pol.MaxRecord]
		rec.Truncated = true
	}

	rec.Data = append([]byte(nil), body...)

	return rec
}

func (r *File) reset() {
	r.buf, r.bufStart, r.dropped, r.started = nil, 0, 0, false
}
