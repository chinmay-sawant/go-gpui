package reader

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/live-log-viewer/entry"
)

func TestStreamLoss(t *testing.T) {
	now := time.Unix(0, 0)
	s := NewDummy(StreamOptions{
		Seed: 1, Key: "api", Start: 1, Rate: 10, Burst: 20,
		Clock: func() time.Time { return now },
		Sleep: func(context.Context, time.Duration) error { return nil },
	})

	ctx := context.Background()

	if b, err := s.Read(ctx); err != nil || len(b.Records) != 0 {
		t.Fatalf("first read = %d records, err=%v", len(b.Records), err)
	}

	now = now.Add(10 * time.Second)

	b, err := s.Read(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(b.Records) != 20 || b.Lost != 80 {
		t.Fatalf("records=%d lost=%d", len(b.Records), b.Lost)
	}
}

func TestBurstCount(t *testing.T) {
	s := NewBurst(StreamOptions{Seed: 1, Start: 1, Count: 3, Burst: 10})
	ctx := context.Background()

	b, err := s.Read(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(b.Records) != 3 || !b.Done || b.State != entry.StateDone {
		t.Fatalf("burst batch = %d records, done=%v", len(b.Records), b.Done)
	}

	b2, err := s.Read(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !b2.Done || len(b2.Records) != 0 {
		t.Fatalf("after done = %+v", b2)
	}
}

func TestStreamCancel(t *testing.T) {
	s := NewDummy(StreamOptions{Seed: 1, Rate: 1000000})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := s.Read(ctx); err == nil {
		t.Fatal("cancelled read returned nil error")
	}
}
