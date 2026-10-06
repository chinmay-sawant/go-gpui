package game

// Result summarizes one completed game for storage. ID is the stable
// unique game ID; re-inserting the same ID must not duplicate a score.
type Result struct {
	ID             string
	Score          int
	Lines          int
	Level          int
	Pieces         int
	DurationMS     int64
	Seed           uint64
	Ruleset        string
	FixtureVersion int
	Dummy          bool
}
