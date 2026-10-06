// Package entry holds the domain types of the live log viewer: severity,
// session and source identity, one stored log entry, a raw record as a
// reader produces it, and the ingestion policy readers and the store share.
package entry

// SessionID, SourceID, and EntryID identify stored rows. Zero means none.
type (
	SessionID int64
	SourceID  int64
	EntryID   int64
)
