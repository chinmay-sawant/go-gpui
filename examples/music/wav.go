package music

import "encoding/binary"

// wavBytes wraps 16-bit stereo PCM in a canonical RIFF/WAVE header.
func wavBytes(pcm []int16) []byte {
	data := make([]byte, 44+len(pcm)*2)
	copy(data, "RIFF")
	binary.LittleEndian.PutUint32(data[4:], uint32(36+len(pcm)*2))
	copy(data[8:], "WAVEfmt ")
	binary.LittleEndian.PutUint32(data[16:], 16)
	binary.LittleEndian.PutUint16(data[20:], 1)
	binary.LittleEndian.PutUint16(data[22:], 2)
	binary.LittleEndian.PutUint32(data[24:], demoRate)
	binary.LittleEndian.PutUint32(data[28:], demoRate*4)
	binary.LittleEndian.PutUint16(data[32:], 4)
	binary.LittleEndian.PutUint16(data[34:], 16)
	copy(data[36:], "data")
	binary.LittleEndian.PutUint32(data[40:], uint32(len(pcm)*2))

	for i, s := range pcm {
		binary.LittleEndian.PutUint16(data[44+i*2:], uint16(s))
	}

	return data
}
