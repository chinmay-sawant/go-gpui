package collector

import (
	"strings"
	"time"
)

// recordError stores the latest failure for one step.
func (m *Manager) recordError(step, source string, err error) {
	if err == nil {
		return
	}

	m.mu.Lock()
	m.errs[step] = Error{Step: step, Source: source, At: time.Now(), Err: strings.TrimSpace(err.Error())}
	switch step {
	case "sample":
		m.sampleErrs++
	case "processes":
		m.procErrs++
	case "detail":
		m.detailErrs++
	}
	m.mu.Unlock()
}

// bumpSkipped counts a tick that was skipped because the previous call of the
// same kind was still running or the pool was full.
func (m *Manager) bumpSkipped() {
	m.mu.Lock()
	m.skipped++
	m.mu.Unlock()
}
