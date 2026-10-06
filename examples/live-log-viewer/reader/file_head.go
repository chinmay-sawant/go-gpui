package reader

import (
	"hash/fnv"
	"io"
	"os"
)

// headBytes bounds the fingerprint taken from the start of a file. It is
// stable across appends because the window length is stored with the hash,
// and it changes when the path is recreated with different content, which
// is what catches inode reuse.
const headBytes = 64

// fileHeadHash fingerprints the first n bytes of an open file. A zero n
// means "pick a window"; the returned length is what the caller must store
// so a later check hashes exactly the same bytes. Zero means the file is
// too short or unreadable, and the caller treats it as no signal.
func fileHeadHash(f *os.File, size, n int64) (uint64, int64) {
	if f == nil || size <= 0 {
		return 0, 0
	}

	if n <= 0 {
		n = size
		if n > headBytes {
			n = headBytes
		}
	}

	if n <= 0 || size < n {
		return 0, 0
	}

	buf := make([]byte, n)

	read, err := f.ReadAt(buf, 0)
	if read <= 0 || (err != nil && err != io.EOF) {
		return 0, 0
	}

	return bytesHash(buf[:read]), int64(read)
}

func bytesHash(data []byte) uint64 {
	h := fnv.New64a()
	_, _ = h.Write(data)

	return h.Sum64()
}

// refreshHead re-fingerprints the open handle after an in-place rewrite.
func (r *File) refreshHead() {
	if r.f == nil {
		return
	}

	head, n := fileHeadHash(r.f, r.size, 0)
	r.head, r.headLen = head, n
	r.opts.HeadHash, r.opts.HeadLen = head, n
}
