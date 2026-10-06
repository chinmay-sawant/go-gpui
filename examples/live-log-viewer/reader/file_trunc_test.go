package reader

import (
	"context"
	"testing"
	"time"
)

func TestFileMaxRecord(t *testing.T) {
	path := logPath(t)
	writeLog(t, path, "1234567890\nshort\n")

	r := testFile(t, path, 0)
	r.pol.MaxRecord = 8

	recs := drain(t, r, 2)

	if !recs[0].Truncated || string(recs[0].Data) != "12345678" {
		t.Fatalf("truncated record = %+v", recs[0])
	}

	if recs[0].Skipped != 2 || recs[0].Bytes != 11 {
		t.Fatalf("skipped=%d bytes=%d", recs[0].Skipped, recs[0].Bytes)
	}

	if string(recs[1].Data) != "short" || recs[1].Offset != 11 {
		t.Fatalf("record after truncation = %+v", recs[1])
	}
}

func TestFileSplitRune(t *testing.T) {
	path := logPath(t)

	writeLog(t, path, "INFO caf\xc3")

	r := testFile(t, path, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	if b, err := r.Read(ctx); err != nil || len(b.Records) != 0 {
		t.Fatalf("early read = %+v, err=%v", b.Records, err)
	}

	appendLog(t, path, []byte("\xa9\n"))

	recs := drain(t, r, 1)
	if string(recs[0].Data) != "INFO café" {
		t.Fatalf("data = %q", recs[0].Data)
	}
}

func TestFileLagNotLoss(t *testing.T) {
	path := logPath(t)
	writeLog(t, path, "aaaaaaaaaa\nbbbbbbbbbb\ncccccccccc\n")

	r := testFile(t, path, 0)
	r.pol.BatchBytes = 12

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	b, err := r.Read(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if b.Lost != 0 {
		t.Fatalf("replayable file reported loss %d", b.Lost)
	}

	if b.Lag <= 0 {
		t.Fatalf("expected lag, got %d", b.Lag)
	}
}
