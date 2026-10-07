package scheduler

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/transfer"
)

// reserveDestinationLocked picks a path no active in-memory job holds.
func (e *Engine) reserveDestinationLocked(dir, name string) (string, error) {
	dest, err := transfer.UniqueDestination(dir, name)
	if err != nil {
		return "", err
	}

	for i := 1; e.destinationTakenLocked(dest) && i < 100; i++ {
		dest, err = transfer.UniqueDestination(dir, transfer.Numbered(name, i))
		if err != nil {
			return "", err
		}
	}

	return dest, nil
}

// destinationTakenLocked reports whether an active job or a reservation
// already owns dest.
func (e *Engine) destinationTakenLocked(dest string) bool {
	if e.reserved[dest] {
		return true
	}

	for _, l := range e.jobs {
		if l.job.Destination == dest && l.job.State.Active() {
			return true
		}
	}

	return false
}

// validateAdd rejects a request that cannot become a job.
func validateAdd(req AddRequest) error {
	if req.Dir == "" {
		return errors.New("scheduler: AddRequest.Dir is required")
	}

	u, err := url.Parse(req.URL)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
		return fmt.Errorf("scheduler: URL %q is not an http or https URL", req.URL)
	}

	if req.Checksum != "" && !strings.HasPrefix(req.Checksum, "sha256:") {
		return fmt.Errorf("scheduler: unsupported checksum %q", req.Checksum)
	}

	return nil
}

// expectedOf maps a zero length to Unknown.
func expectedOf(v int64) int64 {
	if v <= 0 {
		return transfer.Unknown
	}

	return v
}
