package store

// historyQuery builds the keyset query. Terminal states only; the active
// list comes from the scheduler.
func historyQuery(after Cursor, limit int) (string, []any) {
	query := `SELECT ` + jobColumns + ` FROM jobs
		WHERE state IN ('completed','failed','cancelled')`

	args := []any{}

	if !after.UpdatedAt.IsZero() || after.ID != "" {
		query += ` AND (updated_ms < ? OR (updated_ms = ? AND id < ?))`
		at := msOf(after.UpdatedAt)
		args = append(args, at, at, after.ID)
	}

	query += ` ORDER BY updated_ms DESC, id DESC LIMIT ?`
	args = append(args, limit)

	return query, args
}
