package wire

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
	"github.com/chinmay-sawant/ownframe/examples/download-manager/ui"
)

// speed is the smoothed transfer rate of one job, built from successive
// progress events.
type speed struct {
	done  int64
	at    time.Time
	bytes float64
}

// activeRows maps the engine's live list to printable rows.
func (b *Backend) activeRows() []ui.Row {
	jobs := b.eng.Active()
	rows := make([]ui.Row, 0, len(jobs))

	for _, job := range jobs {
		rows = append(rows, b.mapJob(job))
	}

	return rows
}

// rows maps core jobs to printable rows.
func (b *Backend) rows(jobs []domain.Job) []ui.Row {
	rows := make([]ui.Row, 0, len(jobs))

	for _, job := range jobs {
		rows = append(rows, b.mapJob(job))
	}

	return rows
}

// mapJob converts one core job into a printable row, adding the tracked
// speed and the derived ETA.
func (b *Backend) mapJob(job domain.Job) ui.Row {
	row := ui.Row{
		ID:          job.ID,
		Name:        job.Name,
		URL:         job.URL,
		State:       mapState(job.State),
		Done:        job.Done,
		Total:       totalOf(job),
		Attempts:    job.Attempts,
		Destination: job.Destination,
		Error:       job.Error,
		Updated:     job.UpdatedAt,
	}

	b.mu.Lock()
	tracked, ok := b.speeds[job.ID]
	b.mu.Unlock()

	if ok {
		row.Speed = tracked.bytes
	}

	if row.State == ui.StateRunning && row.Total > 0 && row.Speed > 0 {
		row.ETA = time.Duration(float64(row.Total-row.Done) / row.Speed * float64(time.Second))
	}

	return row
}
