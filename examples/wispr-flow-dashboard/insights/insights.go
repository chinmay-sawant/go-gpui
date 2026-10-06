// Package insights is the Wispr Flow insights dashboard: the usage cards,
// the streak heatmap, and the voice and leaderboard tabs. One HTML fragment
// per card, one stylesheet per card, and one Data struct. ownframe renders it.
// This package does not open a window.
package insights

import (
	"embed"
	"strings"
)

//go:embed components/*.html components/*.css
var files embed.FS

// Data is the data the dashboard and its tabs print.
type Data struct {
	ActiveTab   string
	WPM         WPM
	Fixes       Fixes
	Words       Words
	Apps        Apps
	Streak      Streak
	VoiceTab    VoiceTabData
	Leaderboard LeaderboardData
}

// HTML returns the dashboard fragment.
func HTML() string {
	return pageHTML()
}

// CSS returns the dashboard stylesheets in order.
func CSS() string {
	var b strings.Builder

	for _, name := range []string{
		"dashboard.css", "header.css", "tabs.css", "wpm.css", "fixes.css",
		"words.css", "apps.css", "streak.css",
		"insights_voice.css", "leaderboard.css",
	} {
		b.WriteString(file("components/" + name))
	}

	return b.String()
}

// file returns one embedded fragment as text.
func file(name string) string {
	b, err := files.ReadFile(name)
	if err != nil {
		return ""
	}

	return string(b)
}
