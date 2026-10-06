package reader

// readAvailable reads one bounded batch and refreshes the file state.
func (r *File) readAvailable() (Batch, error) {
	b := Batch{Generation: r.gen, Identity: r.identity, Position: r.pos, Size: r.size}

	if err := r.checkPath(&b); err != nil {
		return b, err
	}

	if r.reopened {
		b.Rotated = true
		r.reopened = false
	}

	var bytes int

	for len(b.Records) < r.pol.BatchRecords && bytes < r.pol.BatchBytes {
		chunk, err := r.readChunk(r.pol.BatchBytes - bytes)
		if err != nil {
			return b, err
		}

		if len(chunk) == 0 {
			break
		}

		bytes += len(chunk)

		recs, used := r.consume(chunk, r.pol.BatchRecords-len(b.Records))
		b.Records = append(b.Records, recs...)

		if used < len(chunk) {
			// A record boundary: give the rest of the chunk back.
			r.pos -= int64(len(chunk) - used)

			break
		}
	}

	b.Generation, b.Identity, b.Position, b.Size = r.gen, r.identity, r.pos, r.size
	b.HeadHash, b.HeadLen = r.head, r.headLen
	if r.size > r.pos {
		b.Lag = r.size - r.pos
		b.More = true
	}

	return b, nil
}
