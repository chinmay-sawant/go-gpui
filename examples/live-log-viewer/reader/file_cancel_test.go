package reader

import (
	"context"
	"testing"
)

func TestFileCancel(t *testing.T) {
	path := logPath(t)
	writeLog(t, path, "")

	r := testFile(t, path, 0)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := r.Read(ctx); err == nil {
		t.Fatal("cancelled read returned nil error")
	}
}
