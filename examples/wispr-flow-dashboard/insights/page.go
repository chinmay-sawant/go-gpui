package insights

import "strings"

// pageHTML assembles the dashboard and its three tabs.
func pageHTML() string {
	var b strings.Builder

	b.WriteString(`<div class="page">`)
	b.WriteString(file("components/header.html"))
	b.WriteString(file("components/tabs.html"))
	b.WriteString(`{{if eq .ActiveTab "usage"}}<div class="top-grid">`)
	b.WriteString(file("components/wpm.html"))
	b.WriteString(file("components/fixes.html"))
	b.WriteString(file("components/words.html"))
	b.WriteString(`</div><div class="bottom-grid">`)
	b.WriteString(file("components/apps.html"))
	b.WriteString(file("components/streak.html"))
	b.WriteString(`</div>{{else if eq .ActiveTab "voice"}}`)
	b.WriteString(file("components/insights_voice.html"))
	b.WriteString(`{{else}}`)
	b.WriteString(file("components/leaderboard.html"))
	b.WriteString(`{{end}}</div>`)

	return b.String()
}
