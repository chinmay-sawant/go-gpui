package ui

import "time"

// The bounded budgets the tick and the pager use.
const (
	// drainBudget is the most worker results one tick applies. A deeper
	// queue waits for the next frame, so one busy worker never stalls the
	// UI loop.
	drainBudget = 64

	// summaryEvery and historyEvery bound how often the tick asks for
	// counts and a refreshed history page.
	summaryEvery = 2 * time.Second
	historyEvery = time.Second
)
