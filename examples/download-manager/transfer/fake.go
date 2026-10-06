package transfer

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// FakeOptions configures the dummy transport.
type FakeOptions struct {
	// Dir is the directory the dummy transport writes into. Empty means
	// DummyDir().
	Dir string
	// Pace is the pause between chunks. Zero means 25 ms; tests set a
	// small value.
	Pace time.Duration
	// Chunk is the number of bytes per progress tick. Zero means 32 KiB.
	Chunk int64
}

// DummyDir is the dedicated temporary directory dummy jobs write to.
func DummyDir() string {
	return filepath.Join(os.TempDir(), "ownframe-download-manager-dummy")
}

// Fake is a deterministic transport for dummy mode. It never opens a
// network connection. The body size, progress steps, and one-time failures
// derive from the job ID, so a rerun shows the same behavior.
type Fake struct {
	dir   string
	pace  time.Duration
	chunk int64
}

// NewFake builds the dummy transport.
func NewFake(opts FakeOptions) *Fake {
	if opts.Dir == "" {
		opts.Dir = DummyDir()
	}

	if opts.Pace <= 0 {
		opts.Pace = 25 * time.Millisecond
	}

	if opts.Chunk <= 0 {
		opts.Chunk = 32 << 10
	}

	return &Fake{dir: opts.Dir, pace: opts.Pace, chunk: opts.Chunk}
}

// Download streams the deterministic body to req.Partial and finalizes it.
func (f *Fake) Download(ctx context.Context, req Request, report Reporter) (Outcome, error) {
	if err := os.MkdirAll(filepath.Dir(req.Dest), 0o755); err != nil {
		return Outcome{}, err
	}

	size := FakeSize(req.JobID)
	total := size

	if fakeUnknown(req.JobID) {
		total = Unknown
	}

	offset, resumed, err := partialOffset(req.Partial, size)
	if err != nil {
		return Outcome{}, err
	}

	err = f.stream(ctx, req, offset, size, total, report)
	if err != nil {
		return Outcome{}, err
	}

	if err := CheckChecksum(req.Partial, req.Checksum); err != nil {
		return Outcome{}, err
	}

	if err := Finalize(req.Partial, req.Dest, FinalizeOptions{Sync: true}); err != nil {
		return Outcome{}, err
	}

	return Outcome{
		Path:    req.Dest,
		Bytes:   size,
		Resumed: resumed,
	}, nil
}

// stream writes the body from offset to size, reporting per chunk, and
// fails once mid-body for the jobs fakeFails marks.
func (f *Fake) stream(ctx context.Context, req Request, offset, size, total int64, report Reporter) error {
	stop := int64(-1)

	if fakeFails(req.JobID) && req.Attempt == 0 {
		stop = size * 3 / 5
		if stop <= offset {
			stop = -1
		}
	}

	file, err := os.OpenFile(req.Partial, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Seek(offset, 0); err != nil {
		return err
	}

	done := offset
	buf := make([]byte, f.chunk)
	for done < size {
		if stop >= 0 && done >= stop {
			return fmt.Errorf("transfer: fake connection reset at %d bytes", done)
		}

		if err := sleep(ctx, f.pace); err != nil {
			return err
		}

		n := int64(len(buf))
		if n > size-done {
			n = size - done
		}

		copy(buf, bodyBytes(done, n))
		if _, err := file.Write(buf[:n]); err != nil {
			return err
		}

		done += n
		if report != nil {
			report(Progress{Done: done, Total: total})
		}
	}

	return nil
}

// partialOffset returns the resume offset and whether the transfer resumes.
// A partial larger than the body restarts from zero.
func partialOffset(partial string, size int64) (int64, bool, error) {
	info, err := os.Stat(partial)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, false, nil
		}

		return 0, false, err
	}

	offset := info.Size()
	if offset > size {
		if err := os.Truncate(partial, 0); err != nil {
			return 0, false, err
		}

		return 0, false, nil
	}

	return offset, offset > 0, nil
}

// sleep waits for d or until ctx ends.
func sleep(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return ctx.Err()
	}

	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
