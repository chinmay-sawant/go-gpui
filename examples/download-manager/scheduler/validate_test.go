package scheduler

import (
	"context"
	"testing"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// TestAddValidatesRequest rejects unusable URLs and checksums.
func TestAddValidatesRequest(t *testing.T) {
	eng, _ := newTestEngine(t, instantTransport{}, Options{})
	ctx := context.Background()

	if _, err := eng.Add(ctx, AddRequest{URL: "ftp://x/a", Dir: t.TempDir()}); err == nil {
		t.Error("ftp URL accepted")
	}

	if _, err := eng.Add(ctx, AddRequest{URL: "https://x/a", Dir: ""}); err == nil {
		t.Error("empty dir accepted")
	}

	if _, err := eng.Add(ctx, AddRequest{URL: "https://x/a", Dir: t.TempDir(), Checksum: "md5:1"}); err == nil {
		t.Error("md5 checksum accepted")
	}
}

// testSchedulerJob builds a persisted row for Recover tests.
func testSchedulerJob(id string, state domain.State) domain.Job {
	now := time.Now().UTC().Truncate(time.Millisecond)

	return domain.Job{
		ID:          id,
		URL:         "https://example.invalid/" + id,
		Destination: "/tmp/scheduler-test/" + id,
		Name:        id,
		State:       state,
		Total:       -1,
		Expected:    -1,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
