package transfer

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"
)

// receive writes the response body to the partial file and reports
// progress. The stall guard fails the transfer when no byte arrives within
// the idle window.
func (h *HTTP) receive(ctx context.Context, req Request, resp *http.Response, p plan, report Reporter) (Outcome, error) {
	file, err := openPartial(req.Partial, p)
	if err != nil {
		return Outcome{}, err
	}

	var stalled int32

	guardCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	timer := time.AfterFunc(h.stall, func() {
		atomic.StoreInt32(&stalled, 1)
		cancel()
	})
	defer timer.Stop()

	reset := func() { timer.Reset(h.stall) }

	done, err := pump(guardCtx, resp.Body, file, p, report, &stalled, reset)
	closeErr := file.Close()

	if err != nil {
		return Outcome{}, err
	}

	if closeErr != nil {
		return Outcome{}, closeErr
	}

	if p.total >= 0 && done != p.total {
		return Outcome{}, fmt.Errorf("%w: read %d of %d bytes", ErrBadStatus, done, p.total)
	}

	return h.complete(req, p, Outcome{Bytes: done, Resumed: p.resumed})
}

// pump copies body into file, resetting the stall timer on every read.
func pump(ctx context.Context, body io.Reader, file *os.File, p plan, report Reporter, stalled *int32, reset func()) (int64, error) {
	buf := make([]byte, 64<<10)
	done := p.start

	for {
		n, readErr := body.Read(buf)
		if n > 0 {
			if int64(n) > math.MaxInt64-done {
				return done, fmt.Errorf("transfer: byte counter overflow")
			}

			if _, err := file.Write(buf[:n]); err != nil {
				return done, err
			}

			done += int64(n)
			reset()

			if report != nil {
				report(Progress{Done: done, Total: p.total})
			}
		}

		if readErr == io.EOF {
			return done, nil
		}

		if readErr != nil {
			if atomic.LoadInt32(stalled) == 1 {
				return done, fmt.Errorf("%w: %v", ErrStall, readErr)
			}

			if err := ctx.Err(); err != nil {
				return done, err
			}

			return done, readErr
		}
	}
}

// openPartial opens or creates the partial file at the plan's offset.
func openPartial(path string, p plan) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}

	flags := os.O_CREATE | os.O_WRONLY
	if p.truncate {
		flags |= os.O_TRUNC
	}

	file, err := os.OpenFile(path, flags, 0o644)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrPartialUnwritable, err)
	}

	if _, err := file.Seek(p.start, io.SeekStart); err != nil {
		file.Close()

		return nil, err
	}

	return file, nil
}

// complete checks the checksum, renames the partial into place, and
// reports the outcome. Every handle must already be closed.
func (h *HTTP) complete(req Request, p plan, out Outcome) (Outcome, error) {
	if err := CheckChecksum(req.Partial, req.Checksum); err != nil {
		return Outcome{}, err
	}

	if err := Finalize(req.Partial, req.Dest, FinalizeOptions{Sync: true}); err != nil {
		return Outcome{}, err
	}

	size, err := statSize(req.Dest)
	if err != nil {
		return Outcome{}, err
	}

	out.Path = req.Dest
	out.Bytes = size
	out.Total = p.total
	out.Validators = p.validators

	return out, nil
}
