package reader

// readAvailable reads one bounded batch and refreshes the file state.
func (r *File) readAvailable() (Batch, error) {
	b := Batch{Generation: r.gen, Identity: r.identity, Position: r.pos, Size: r.size}

	if err := r.checkPath(&b); err != nil {
		return b, err
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
		b.Records = append(b.Records, r.consume(chunk)...)
	}

	b.Generation, b.Identity, b.Position, b.Size = r.gen, r.identity, r.pos, r.size
	if r.size > r.pos {
		b.Lag = r.size - r.pos
		b.More = true
	}

	return b, nil
}
