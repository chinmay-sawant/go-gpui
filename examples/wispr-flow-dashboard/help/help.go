// Package help is the help page: search, topics, and popular articles.
package help

// Data is the data the help page prints.
type Data struct {
	Topics   []HelpTopic
	Articles []string
}

// HelpTopic is one card in the help topic grid.
type HelpTopic struct {
	Icon  string
	Title string
	Line  string
}

// Default returns the sample help topics and popular articles.
func Default() Data {
	return Data{
		Topics: []HelpTopic{
			{"icon-play", "Getting started", "Install Flow and dictate your first sentence."},
			{"icon-mic", "Dictation basics", "Start, stop, and format dictation in any app."},
			{"icon-note", "Notetaker", "Capture meetings and conversations automatically."},
			{"icon-wand", "Troubleshooting", "Fix microphone, permission, and typing issues."},
			{"icon-text", "Keyboard shortcuts", "See every hotkey Flow listens for."},
			{"icon-gear", "Account & billing", "Manage your plan, invoices, and receipts."},
		},
		Articles: []string{
			"Voice commands cheat sheet",
			"Flow is not typing into my app",
			"Fix repeated words in dictations",
			"Set up a second language",
			"Export your dictation history",
		},
	}
}
