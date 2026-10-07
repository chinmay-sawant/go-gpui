package scheduler

import (
	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// request builds the transport request for a job. Attempt is zero-based,
// so a fake can fail only the first try.
func (e *Engine) request(job domain.Job) transfer.Request {
	attempt := job.Attempts - 1
	if attempt < 0 {
		attempt = 0
	}

	return transfer.Request{
		JobID:      job.ID,
		URL:        job.URL,
		Dest:       job.Destination,
		Partial:    transfer.PartialPath(job.Destination),
		Expected:   job.Expected,
		Validators: transfer.Validators{ETag: job.ETag, LastModified: job.LastModified},
		Checksum:   job.Checksum,
		Attempt:    attempt,
	}
}
