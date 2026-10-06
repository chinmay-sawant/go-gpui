package reader

import (
	"context"
	"strings"
	"testing"
	"time"
)

func TestFileAppendCRLF(t *testing.T) {
	path := logPath(t)
	writeLog(t, path, "INFO one\r\nWARN two\n")

	r := testFile(t, path, 0)
	recs := drain(t, r, 2)

	if string(recs[0].Data) != "INFO one" || recs[0].Offset != 0 || recs[0].Bytes != 10 {
		t.Fatalf("record 0 = %+v", recs[0])
	}

	if string(recs[1].Data) != "WARN two" || recs[1].Offset != 10 || recs[1].Bytes != 9 {
		t.Fatalf("record 1 = %+v", recs[1])
	}

	appendLog(t, path, []byte("ERROR three\n"))

	more := drain(t, r, 1)
	if string(more[0].Data) != "ERROR three" || more[0].Offset != 19 {
		t.Fatalf("appended record = %+v", more[0])
	}
}

func TestFileBatchCap(t *testing.T) {
	path := logPath(t)
	writeLog(t, path, strings.Repeat("x\n", 10))

	r := testFile(t, path, 0)
	r.pol.BatchRecords = 3

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	b, err := r.Read(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(b.Records) != 3 || !b.More {
		t.Fatalf("batch = %d records, more=%v", len(b.Records), b.More)
	}

	if b.Position != 6 {
		t.Fatalf("position = %d, want 6", b.Position)
	}

	recs := drain(t, r, 7)
	if len(recs) != 7 {
		t.Fatalf("drained %d records", len(recs))
	}
}

func TestFilePartialLine(t *testing.T) {
	path := logPath(t)
	writeLog(t, path, "INFO done\nINFO partial")

	r := testFile(t, path, 0)
	drain(t, r, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	b, err := r.Read(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(b.Records) != 0 {
		t.Fatalf("unterminated line emitted: %+v", b.Records)
	}

	part := r.Flush()
	if len(part) != 1 || !part[0].Partial {
		t.Fatalf("flush = %+v", part)
	}

	if string(part[0].Data) != "INFO partial" || part[0].Offset != 10 || part[0].Bytes != 12 {
		t.Fatalf("partial record = %+v", part[0])
	}
}
