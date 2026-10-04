// Package style is the style page: writing styles and per-app choices.
package style

// StyleCard is one selectable writing style.
type StyleCard struct {
	Name     string
	Text     string
	Selected bool
}

// StyleApp is one app and the style Flow uses for it.
type StyleApp struct {
	App   string
	Style string
}

// Data is the data the style page prints.
type Data struct {
	Cards []StyleCard
	Apps  []StyleApp
}

// Default returns the style page data.
func Default() Data {
	return Data{
		Cards: []StyleCard{
			{Name: "Formal", Text: "Polished, complete sentences."},
			{Name: "Casual", Text: "Relaxed, everyday language.", Selected: true},
			{Name: "Concise", Text: "Short, direct answers."},
			{Name: "Email", Text: "Greetings, sign-offs, and structure."},
		},
		Apps: []StyleApp{
			{App: "Gmail", Style: "Formal"},
			{App: "Slack", Style: "Casual"},
			{App: "Notes", Style: "Concise"},
			{App: "Scratchpad", Style: "Casual"},
		},
	}
}
