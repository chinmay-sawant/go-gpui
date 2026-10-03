package music

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestDemoTuneIsCanonicalWAV(t *testing.T) {
	t.Parallel()

	clip := DemoTune("jazz")
	if !bytes.HasPrefix(clip.Data, []byte("RIFF")) {
		t.Fatal("demo data is not a RIFF file")
	}

	want := 44 + demoSeconds*demoRate*4
	if len(clip.Data) != want {
		t.Fatalf("len = %d, want %d", len(clip.Data), want)
	}

	rate := binary.LittleEndian.Uint32(clip.Data[24:])
	if rate != demoRate {
		t.Fatalf("sample rate = %d", rate)
	}

	if clip.License != "generated" || clip.Creator == "" {
		t.Fatalf("clip = %+v", clip)
	}
}

func TestDemoTuneIsDeterministic(t *testing.T) {
	t.Parallel()

	first := DemoTune("jazz")
	second := DemoTune("jazz")

	if !bytes.Equal(first.Data, second.Data) {
		t.Fatal("the same query produced different audio")
	}

	if bytes.Equal(first.Data, DemoTune("rock").Data) {
		t.Fatal("different queries produced the same audio")
	}
}
