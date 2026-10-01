//go:build linux && !android

package clipboard

func internReq(name string) []byte {
	size := pad4(8 + len(name))
	buf := make([]byte, size)
	buf[0] = 16
	put16(buf[2:], uint16(size/4))
	put16(buf[4:], uint16(len(name)))
	copy(buf[8:], name)

	return buf
}

func winReq(id, parent uint32) []byte {
	buf := make([]byte, 32)
	buf[0] = 1
	put16(buf[2:], 8)
	put32(buf[4:], id)
	put32(buf[8:], parent)
	put16(buf[16:], 1)
	put16(buf[18:], 1)
	put16(buf[22:], 1)

	return buf
}

func ownReq(owner, sel uint32) []byte {
	buf := make([]byte, 16)
	buf[0] = 22
	put16(buf[2:], 4)
	put32(buf[4:], owner)
	put32(buf[8:], sel)

	return buf
}

func convReq(win, sel, target, prop uint32) []byte {
	buf := make([]byte, 24)
	buf[0] = 24
	put16(buf[2:], 6)
	put32(buf[4:], win)
	put32(buf[8:], sel)
	put32(buf[12:], target)
	put32(buf[16:], prop)

	return buf
}

func ownerReq(sel uint32) []byte {
	buf := make([]byte, 8)
	buf[0] = 23
	put16(buf[2:], 2)
	put32(buf[4:], sel)

	return buf
}

func getReq(win, prop uint32) []byte {
	buf := make([]byte, 24)
	buf[0] = 20
	buf[1] = 1
	put16(buf[2:], 6)
	put32(buf[4:], win)
	put32(buf[8:], prop)
	put32(buf[20:], 1<<20)

	return buf
}

func prop8(win, prop, typ uint32, data []byte) []byte {
	size := pad4(24 + len(data))
	buf := make([]byte, size)
	buf[0] = 18
	put16(buf[2:], uint16(size/4))
	put32(buf[4:], win)
	put32(buf[8:], prop)
	put32(buf[12:], typ)
	buf[16] = 8
	put32(buf[20:], uint32(len(data)))
	copy(buf[24:], data)

	return buf
}
