package insights

// LeaderboardData is the data the insights leaderboard tab prints.
type LeaderboardData struct {
	Rows []LeaderRow
}

// LeaderRow is one friend on the leaderboard. You marks the viewer's row.
type LeaderRow struct {
	Rank   string
	Name   string
	Words  string
	Streak string
	You    bool
}

// DefaultLeaderboard returns the sample weekly leaderboard.
func DefaultLeaderboard() LeaderboardData {
	return LeaderboardData{Rows: []LeaderRow{
		{"1", "Priya", "41,208", "63", false},
		{"2", "Arjun", "38,940", "51", false},
		{"3", "Maya", "33,115", "44", false},
		{"4", "You", "29,804", "38", true},
		{"5", "Leo", "24,670", "31", false},
		{"6", "Sara", "21,553", "27", false},
		{"7", "Dev", "18,209", "19", false},
		{"8", "Nina", "12,844", "12", false},
	}}
}
