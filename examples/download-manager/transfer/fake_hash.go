package transfer

import (
	"crypto/sha256"
	"hash"
)

// newSHA256 is a seam so FakeChecksum reads cleanly.
func newSHA256() hash.Hash { return sha256.New() }

// fakeWriteHash feeds the whole deterministic body into h in chunks.
func fakeWriteHash(h hash.Hash, size int64) error {
	const step = 64 << 10
	for off := int64(0); off < size; off += step {
		n := int64(step)
		if n > size-off {
			n = size - off
		}

		if _, err := h.Write(bodyBytes(off, n)); err != nil {
			return err
		}
	}

	return nil
}
