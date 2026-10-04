package app

import "github.com/chinmay-sawant/go-gpui/examples/wispr-flow-dashboard/insights"

// shiftStreak scrolls the heatmap by step weeks; positive steps into the
// past. The offset stops at the oldest and newest windows.
func (a *App) shiftStreak(step int) {
	a.streak = min(max(a.streak+step, 0), insights.MaxStreakOffset())
	a.view.Streak = insights.BuildStreak(a.streak)
}
