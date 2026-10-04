package insights

import "html/template"

// View is the data the dashboard template prints.
type View struct {
	ActiveTab  string
	WPM        WPM
	Fixes      Fixes
	Words      Words
	Apps       Apps
	Streak     Streak
	EmptyTitle string
	EmptyText  string
	Note       string
}

// WPM is the words-per-minute card.
type WPM struct {
	Value string
	Top   string
}

// Fixes is the fixes-made-by-flow card.
type Fixes struct {
	Value      string
	Corrected  string
	Dictionary string
}

// Words is the total-words-dictated card.
type Words struct {
	Value   string
	Delta   string
	Desktop string
}

// Apps is the desktop-usage card.
type Apps struct {
	Total string
	Rows  []AppRow
}

// AppRow is one desktop-usage row. A Bar row draws a full bar instead of a
// percent badge. Color is a pct class: pct-dark, pct-mid, or pct-light.
type AppRow struct {
	Icon    string
	Percent string
	Label   string
	Bar     bool
	Color   string
}

// Streak is the day-streak heatmap card. Prev and Next say whether a
// chevron can scroll to older or newer weeks.
type Streak struct {
	Days    string
	Longest string
	Months  []Month
	Weeks   [][]Cell
	Prev    bool
	Next    bool
}

// Month is one month label over the heatmap.
type Month struct {
	Label  string
	Column template.CSS
}

// Cell is one heatmap day. Level 0 is empty and 5 is the darkest teal.
// Current marks a day inside the current streak.
type Cell struct {
	Level   int
	Current bool
}
