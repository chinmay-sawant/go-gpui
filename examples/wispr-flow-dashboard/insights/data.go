package insights

// DefaultView returns the sample numbers the dashboard draws.
func DefaultView() View {
	return View{
		ActiveTab:  "usage",
		WPM:        WPM{Value: "148", Top: "0.2%"},
		Fixes:      Fixes{Value: "24,882", Corrected: "17,737", Dictionary: "7,145"},
		Words:      Words{Value: "164,134", Delta: "338% this month", Desktop: "164,134 words"},
		Apps:       Apps{Total: "29", Rows: appRows()},
		Streak:     buildStreak(0),
		EmptyTitle: "Your voice",
		EmptyText:  "Nothing to show on this tab yet.",
	}
}

// appRows is the desktop-usage list. The first row draws the bar.
func appRows() []AppRow {
	return []AppRow{
		{Percent: "91%", Label: "3,406 other tasks", Bar: true},
		{Icon: "icon-browser", Percent: "5%", Label: "220 AI prompts", Color: "pct-dark"},
		{Icon: "icon-doc", Percent: "2%", Label: "8 documents", Color: "pct-dark"},
		{Icon: "icon-chat", Percent: "1%", Label: "58 personal messages", Color: "pct-light"},
		{Icon: "icon-mail", Percent: "0%", Label: "4 emails", Color: "pct-light"},
		{Icon: "icon-work", Percent: "0%", Label: "0 work messages", Color: "pct-light"},
	}
}
