package page

import "testing"

func TestTouchScrollOptionsDefaultAndClamp(t *testing.T) {
	p := &Page{}
	if got := p.TouchScrollOptions(); got.Sensitivity != 1 || got.Deceleration != 5 {
		t.Fatalf("default options = %+v", got)
	}
	p.SetTouchScrollOptions(TouchScrollOptions{Sensitivity: .1, Deceleration: 20})
	if got := p.TouchScrollOptions(); got.Sensitivity != .25 || got.Deceleration != 12 {
		t.Fatalf("clamped options = %+v", got)
	}
}
