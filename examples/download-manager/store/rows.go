package store

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/download-manager/domain"
)

// timeNow is a seam tests replace.
var timeNow = time.Now

// msOf converts a time to UTC unix milliseconds, the storage clock.
func msOf(t time.Time) int64 { return t.UTC().UnixMilli() }

// fromMS converts storage milliseconds back to UTC.
func fromMS(ms int64) time.Time { return time.UnixMilli(ms).UTC() }

// jobColumns is the fixed column order every job query shares.
const jobColumns = `id, url, destination, name, state, done, total, expected, ` +
	`etag, last_modified, checksum, attempts, error, created_ms, updated_ms`

// scanner is satisfied by *sql.Row and *sql.Rows.
type scanner interface {
	Scan(dest ...any) error
}

// scanJob reads one row in jobColumns order.
func scanJob(row scanner) (domain.Job, error) {
	var (
		j               domain.Job
		state           string
		created, update int64
	)

	err := row.Scan(&j.ID, &j.URL, &j.Destination, &j.Name, &state,
		&j.Done, &j.Total, &j.Expected, &j.ETag, &j.LastModified,
		&j.Checksum, &j.Attempts, &j.Error, &created, &update)
	if err != nil {
		return domain.Job{}, err
	}

	j.State = domain.State(state)
	j.CreatedAt = fromMS(created)
	j.UpdatedAt = fromMS(update)

	return j, nil
}

// jobArgs flattens a job for the insert statement.
func jobArgs(j domain.Job) []any {
	return []any{
		j.ID, j.URL, j.Destination, j.Name, string(j.State),
		j.Done, j.Total, j.Expected, j.ETag, j.LastModified,
		j.Checksum, j.Attempts, j.Error,
		msOf(j.CreatedAt), msOf(j.UpdatedAt),
	}
}
