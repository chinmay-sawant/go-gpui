package dino

import (
	"testing"
	"time"
)

func TestMeterCountsFrames(t *testing.T) {
	var m meter
	base := time.Unix(0, 0)
	m.add(base)

	for i := 1; i <= 60; i++ {
		m.add(base.Add(time.Duration(i) * time.Second / 60))
	}

	if m.fps != 60 {
		t.Fatalf("fps = %d, want 60", m.fps)
	}

	if m.frames != 0 {
		t.Fatalf("frames = %d, want the counter reset", m.frames)
	}
}

func TestMeterStaysQuietBeforeHalfASecond(t *testing.T) {
	var m meter
	base := time.Unix(0, 0)
	m.add(base)

	m.add(base.Add(100 * time.Millisecond))

	if m.fps != 0 {
		t.Fatalf("fps = %d early", m.fps)
	}
}
