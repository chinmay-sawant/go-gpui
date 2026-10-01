package clipboard

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	UseMemory(true)
	os.Exit(m.Run())
}

func TestMemoryRoundTrip(t *testing.T) {
	UseMemory(true)

	Write("alpha")

	if got := Read(); got != "alpha" {
		t.Fatalf("Read() = %q", got)
	}

	Write("beta\nline")

	if got := Read(); got != "beta\nline" {
		t.Fatalf("Read() = %q", got)
	}
}

func TestUTF16Clipboard(t *testing.T) {
	got := utf16Clipboard("A€")
	want := []byte{
		'A', 0,
		0xac, 0x20,
		0, 0,
	}

	if len(got) != len(want) {
		t.Fatalf("len %d", len(got))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("byte %d = %d", i, got[i])
		}
	}
}
