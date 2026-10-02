package replay

import "testing"

func TestImageBytesKeyStableAndDistinct(t *testing.T) {
	data := tinyPNG(t, 255)

	first := imageOp(t, data, 2, 2, false)
	same := imageOp(t, append([]byte(nil), data...), 2, 2, false)
	if imageBytesKey(&first) != imageBytesKey(&same) {
		t.Fatal("same payload produced different keys")
	}

	other := imageOp(t, tinyPNG(t, 7), 2, 2, false)
	if imageBytesKey(&first) == imageBytesKey(&other) {
		t.Fatal("different payloads share a key")
	}

	jpeg := imageOp(t, data, 2, 2, true)
	if imageBytesKey(&first) == imageBytesKey(&jpeg) {
		t.Fatal("IsJPEG is not folded into the key")
	}

	if imageBytesKey(nil) != 0 {
		t.Fatal("nil op key is not zero")
	}
}
