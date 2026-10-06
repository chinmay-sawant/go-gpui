package transfer

import (
	"context"
	"fmt"
	"os"
	"testing"
)

// findID returns the first ID whose predicate matches.
func findID(pred func(string) bool) string {
	for i := 0; ; i++ {
		id := fmt.Sprintf("probe-%d", i)
		if pred(id) {
			return id
		}
	}
}

// TestFakeFailsFirstAttempt drops the first try and succeeds on retry.
func TestFakeFailsFirstAttempt(t *testing.T) {
	dest, partial := fakePaths(t)
	id := findID(fakeFails)

	_, err := fastFake().Download(context.Background(), Request{
		JobID: id, Dest: dest, Partial: partial, Attempt: 0,
	}, nil)
	if err == nil {
		t.Fatal("first attempt succeeded")
	}

	if _, statErr := os.Stat(partial); statErr != nil {
		t.Fatal("partial not kept for resume")
	}

	out, err := fastFake().Download(context.Background(), Request{
		JobID: id, Dest: dest, Partial: partial, Attempt: 1,
	}, nil)
	if err != nil {
		t.Fatalf("retry: %v", err)
	}

	if !out.Resumed {
		t.Error("retry did not resume")
	}
}

// TestFakeUnknownLength reports Total as Unknown.
func TestFakeUnknownLength(t *testing.T) {
	dest, partial := fakePaths(t)
	id := findID(fakeUnknown)

	out, err := fastFake().Download(context.Background(), Request{
		JobID: id, Dest: dest, Partial: partial,
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	if out.Total != Unknown {
		t.Errorf("total %d, want Unknown", out.Total)
	}

	if out.Bytes != FakeSize(id) {
		t.Errorf("bytes %d, want %d", out.Bytes, FakeSize(id))
	}
}
