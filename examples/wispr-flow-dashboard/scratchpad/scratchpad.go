// Package scratchpad is the scratchpad page: a note area and recent pads.
package scratchpad

// Data is the data the scratchpad page prints.
type Data struct {
	Note     string
	WordLine string
	Recent   []RecentNote
}

// RecentNote is one row in the recent scratchpads list.
type RecentNote struct {
	Title string
	When  string
}

// Default returns the scratchpad page data.
func Default() Data {
	return Data{
		Note: "Quick note before standup: the transforms page is done, the sidebar " +
			"collapse needs a second look, and I still owe Priya a review of the release " +
			"checklist. After standup I will fix the hover state and draft the invite copy.",
		WordLine: "142 words · saved 2 minutes ago",
		Recent: []RecentNote{
			{"Standup notes", "2 minutes ago"},
			{"Release checklist", "Yesterday"},
			{"Interview questions", "Mon"},
			{"Gift ideas", "Last week"},
		},
	}
}
