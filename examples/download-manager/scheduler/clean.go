package scheduler

import (
	"strings"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// cleanError redacts the job URL from a transport error before it reaches
// the UI or the log.
func cleanError(job domain.Job, err error) string {
	text := err.Error()
	if job.URL != "" {
		text = strings.ReplaceAll(text, job.URL, domain.RedactURL(job.URL))
	}

	return text
}

// joinNote appends b to a unless a is empty.
func joinNote(a, b string) string {
	if a == "" {
		return b
	}

	return a + "; " + b
}
