package ui

import "time"

// Row is one queue or history line. Done and Total are bytes; a Total of
// zero or less means the length is unknown.
type Row struct {
	ID          string
	Name        string
	URL         string
	State       State
	Done        int64
	Total       int64
	Speed       float64
	ETA         time.Duration
	Attempts    int
	Destination string
	Error       string
	Updated     time.Time
}

// SpeedText is the transfer rate while the job runs.
func (r Row) SpeedText() string {
	if r.State != StateRunning {
		return "--"
	}

	return formatSpeed(r.Speed)
}

// ETAText is the remaining time while the job runs.
func (r Row) ETAText() string {
	if r.State != StateRunning {
		return "--"
	}

	return formatETA(r.ETA)
}

// SizeText is the observed over expected size.
func (r Row) SizeText() string {
	return formatBytes(r.Done) + " / " + formatBytes(r.Total)
}

// StateText is the symbol and the name, readable without color.
func (r Row) StateText() string {
	return r.State.Symbol() + " " + r.State.Label()
}

// URLText is the URL with any credentials removed.
func (r Row) URLText() string {
	return redactURL(r.URL)
}

// UpdatedText is the last transition time, or "--".
func (r Row) UpdatedText() string {
	if r.Updated.IsZero() {
		return "--"
	}

	return r.Updated.Local().Format("2006-01-02 15:04:05")
}
