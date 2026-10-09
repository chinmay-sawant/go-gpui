package ws

import (
	"crypto/rand"
	"encoding/binary"
	"io"
)

func writeFrame(w io.Writer, op byte, p []byte, mask bool) error {
	hdr := []byte{0x80 | op}
	n := len(p)

	switch {
	case n < 126:
		hdr = append(hdr, byte(n))
	case n < 65536:
		hdr = append(hdr, 126, byte(n>>8), byte(n))
	default:
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], uint64(n))
		hdr = append(hdr, 127)
		hdr = append(hdr, b[:]...)
	}

	var m [4]byte
	if mask {
		hdr[1] |= 0x80
		_, _ = rand.Read(m[:])
		hdr = append(hdr, m[:]...)
	}

	if _, err := w.Write(hdr); err != nil {
		return err
	}

	if !mask {
		_, err := w.Write(p)
		return err
	}

	out := make([]byte, len(p))
	for i := range p {
		out[i] = p[i] ^ m[i%4]
	}

	_, err := w.Write(out)

	return err
}

func readFrame(r io.Reader) (bool, byte, []byte, error) {
	var h [2]byte
	if _, err := io.ReadFull(r, h[:]); err != nil {
		return false, 0, nil, err
	}

	n, err := frameLen(r, int(h[1]&0x7f))
	if err != nil {
		return false, 0, nil, err
	}

	var mask [4]byte
	masked := h[1]&0x80 != 0
	if masked {
		if _, err = io.ReadFull(r, mask[:]); err != nil {
			return false, 0, nil, err
		}
	}

	p := make([]byte, n)
	if _, err = io.ReadFull(r, p); err != nil {
		return false, 0, nil, err
	}

	if masked {
		for i := range p {
			p[i] ^= mask[i%4]
		}
	}

	return h[0]&0x80 != 0, h[0] & 0x0f, p, nil
}

func frameLen(r io.Reader, n int) (int, error) {
	switch n {
	case 126:
		var b [2]byte
		if _, err := io.ReadFull(r, b[:]); err != nil {
			return 0, err
		}

		return int(binary.BigEndian.Uint16(b[:])), nil
	case 127:
		var b [8]byte
		if _, err := io.ReadFull(r, b[:]); err != nil {
			return 0, err
		}

		return int(binary.BigEndian.Uint64(b[:])), nil
	default:
		return n, nil
	}
}
