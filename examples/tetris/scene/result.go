package scene

import "github.com/chinmay-sawant/ownframe/examples/tetris/game"

// ResultKind names one finished store request.
type ResultKind uint8

// The result kinds.
const (
	ResultSettings ResultKind = iota
	ResultTheme
	ResultScores
	ResultDummy
	ResultSaved
	ResultSnapshot
)

// ScoreEntry is one row of the score history.
type ScoreEntry struct {
	Rank  int
	ID    string
	Score int
	Lines int
	Level int
	Dummy bool
}

// ScorePage is one page of the score history, indexed from zero. More is
// true when the store may hold entries after this page.
type ScorePage struct {
	Index   int
	Total   int
	More    bool
	Entries []ScoreEntry
}

// SettingsResult is the answer to Store.Settings.
type SettingsResult struct {
	Settings Settings
	Snapshot game.Snapshot
	Has      bool
}

// Result is one finished store request. ID matches the request that
// produced it, so a stale answer is recognizable.
type Result struct {
	ID     uint64
	Kind   ResultKind
	Set    SettingsResult
	Scores ScorePage
	Err    error
}
