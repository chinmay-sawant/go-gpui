//go:build linux && !android

package clipboard

func prop32(win, prop, typ uint32, vals []uint32) []byte {
	size := 24 + 4*len(vals)
	buf := make([]byte, size)
	buf[0] = 18
	put16(buf[2:], uint16(size/4))
	put32(buf[4:], win)
	put32(buf[8:], prop)
	put32(buf[12:], typ)
	buf[16] = 32
	put32(buf[20:], uint32(len(vals)))

	for i, v := range vals {
		put32(buf[24+4*i:], v)
	}

	return buf
}

func selEvent(when, who, sel, target, prop uint32) []byte {
	ev := make([]byte, 32)
	ev[0] = 31
	put32(ev[4:], when)
	put32(ev[8:], who)
	put32(ev[12:], sel)
	put32(ev[16:], target)
	put32(ev[20:], prop)

	return ev
}

func sendEv(dst uint32, ev []byte) []byte {
	buf := make([]byte, 44)
	buf[0] = 25
	put16(buf[2:], 11)
	put32(buf[4:], dst)
	copy(buf[12:], ev[:32])

	return buf
}
