package transfer

import (
	"context"
	"errors"
	"os"
	"testing"
)

// TestFakeChecksum covers both a match and a mismatch.
func TestFakeChecksum(t *testing.T) {
	dest, partial := fakePaths(t)
	id := "checksum-1"

	_, err := fastFake().Download(context.Background(), Request{
		JobID: id, Dest: dest, Partial: partial, Checksum: FakeChecksum(id),
	}, nil)
	if err != nil {
		t.Fatalf("good checksum: %v", err)
	}

	dest2, partial2 := fakePaths(t)

	_, err = fastFake().Download(context.Background(), Request{
		JobID: id, Dest: dest2, Partial: partial2, Checksum: "sha256:00",
	}, nil)
	if !errors.Is(err, ErrChecksum) {
		t.Fatalf("want ErrChecksum, got %v", err)
	}

	if _, statErr := os.Stat(partial2); !os.IsNotExist(statErr) {
		t.Error("corrupt partial kept")
	}
}
