package store

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetryBusyGivesUpAtTheDeadline(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	calls := 0

	err := retryBusy(ctx, func() error {
		calls++

		return errors.New("database is locked")
	})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("retry returned %v", err)
	}

	if calls < 2 {
		t.Fatalf("retry tried %d times", calls)
	}
}

func TestRetryBusyStopsOnOtherErrors(t *testing.T) {
	calls := 0

	err := retryBusy(context.Background(), func() error {
		calls++

		return errors.New("boom")
	})
	if err == nil || calls != 1 {
		t.Fatalf("retry returned %v after %d calls", err, calls)
	}
}

func TestRetryBusyRecoversFromATemporaryLock(t *testing.T) {
	calls := 0

	err := retryBusy(context.Background(), func() error {
		calls++
		if calls < 3 {
			return errors.New("SQLITE_BUSY")
		}

		return nil
	})
	if err != nil || calls != 3 {
		t.Fatalf("retry returned %v after %d calls", err, calls)
	}
}
