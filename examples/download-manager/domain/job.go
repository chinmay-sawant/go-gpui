package domain

import (
	"time"
)

// Job is one download task as the UI, the store, and the scheduler see it.
// Done is the byte count observed on disk. Total and Expected are -1 when
// the server sent no usable length.
type Job struct {
	ID           string    `json:"id"`
	URL          string    `json:"url"`
	Destination  string    `json:"destination"`
	Name         string    `json:"name"`
	State        State     `json:"state"`
	Done         int64     `json:"done"`
	Total        int64     `json:"total"`
	Expected     int64     `json:"expected"`
	ETag         string    `json:"etag"`
	LastModified string    `json:"lastModified"`
	Checksum     string    `json:"checksum"`
	Attempts     int       `json:"attempts"`
	Error        string    `json:"error"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// Fraction returns Done over the expected length in 0..1. It returns -1
// when the length is unknown and clamps a longer-than-expected file to 1.
func (j Job) Fraction() float64 {
	total := j.Expected
	if total <= 0 {
		total = j.Total
	}

	if total <= 0 || j.Done < 0 {
		return -1
	}

	if j.Done >= total {
		return 1
	}

	return float64(j.Done) / float64(total)
}

// UnknownLength reports whether no usable total is known yet.
func (j Job) UnknownLength() bool {
	return j.Expected <= 0 && j.Total <= 0
}
