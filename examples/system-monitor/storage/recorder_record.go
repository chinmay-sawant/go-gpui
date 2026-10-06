package storage

import "github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"

// RecordSample converts one sample to rows and enqueues them. It reports
// whether the sample was queued.
func (r *Recorder) RecordSample(s domain.Sample) bool { return r.Record(s.Records()) }

// RecorderStats reports what the recorder wrote, dropped, and failed.
type RecorderStats struct {
	Written   int64
	Failed    int64
	Dropped   uint64
	Queued    int
	LastError string
}

// Stats returns a point-in-time view.
func (r *Recorder) Stats() RecorderStats {
	r.mu.Lock()
	defer r.mu.Unlock()

	return RecorderStats{
		Written:   r.written,
		Failed:    r.failed,
		Dropped:   r.dropped.Load(),
		Queued:    len(r.queue),
		LastError: r.lastError,
	}
}

// Record enqueues prepared rows. A full queue drops the rows and counts the
// drop; nothing blocks the caller.
func (r *Recorder) Record(rows []Row) bool {
	if len(rows) == 0 || r.closed.Load() {
		return false
	}

	select {
	case r.queue <- rows:
		return true
	default:
		r.dropped.Add(1)

		return false
	}
}
