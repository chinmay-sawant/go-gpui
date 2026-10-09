package window

import "testing"

func TestFitBoxShowsTheWholePage(t *testing.T) {
	scale, ox, oy := fitBox(200, 400, 100, 100)
	if scale != 0.25 || ox != 25 || oy != 0 {
		t.Fatalf("scale=%v ox=%v oy=%v", scale, ox, oy)
	}
}

func TestFitBoxLeavesAMatchingPage(t *testing.T) {
	scale, ox, oy := fitBox(1080, 2400, 1080, 2400)
	if scale != 1 || ox != 0 || oy != 0 {
		t.Fatalf("scale=%v ox=%v oy=%v", scale, ox, oy)
	}
}
