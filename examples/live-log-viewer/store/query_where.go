package store

import (
	"context"
	"database/sql"
	"strings"
)

func countTx(ctx context.Context, db *sql.DB, q Query) (int64, error) {
	cond, args := whereSQL(q)

	if q.MaxID > 0 {
		cond += " AND id <= ?"
		args = append(args, int64(q.MaxID))
	}

	if q.Cursor > 0 {
		cond += " AND id > ?"
		args = append(args, int64(q.Cursor))
	}

	var n int64

	err := db.QueryRowContext(ctx, `SELECT count(*) FROM entries`+cond, args...).Scan(&n)

	return n, err
}

func whereSQL(q Query) (string, []any) {
	var (
		b    strings.Builder
		args []any
	)

	b.WriteString(" WHERE 1 = 1")

	if q.Session != nil {
		b.WriteString(" AND session_id = ?")
		args = append(args, int64(*q.Session))
	}

	if q.Source != nil {
		b.WriteString(" AND source_id = ?")
		args = append(args, int64(*q.Source))
	}

	if q.MinSeverity != nil {
		b.WriteString(" AND severity >= ?")
		args = append(args, int64(*q.MinSeverity))
	}

	if q.Text != "" {
		b.WriteString(` AND message LIKE '%' || ? || '%' ESCAPE '\'`)
		args = append(args, escapeLike(q.Text))
	}

	if q.Since != nil {
		b.WriteString(" AND received_ns >= ?")
		args = append(args, q.Since.UnixNano())
	}

	if q.Until != nil {
		b.WriteString(" AND received_ns < ?")
		args = append(args, q.Until.UnixNano())
	}

	return b.String(), args
}

func escapeLike(s string) string {
	return strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`).Replace(s)
}
