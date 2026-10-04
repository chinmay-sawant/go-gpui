// Package notetaker is the notetaker page: meeting notes and voice memos.
package notetaker

// NoteRow is one row in the notetaker list.
type NoteRow struct {
	ID    string
	Title string
	Meta  string
	Tag   string
}

// Data is the data the notetaker page prints.
type Data struct {
	Notes []NoteRow
}

// Default returns the notetaker page data.
func Default() Data {
	return Data{
		Notes: []NoteRow{
			{"1", "Weekly product sync", "Today · 42 min", "Meeting"},
			{"2", "Design review with Maya", "Yesterday · 28 min", "Meeting"},
			{"3", "Grocery list", "Yesterday · 2 min", "Voice memo"},
			{"4", "Interview notes — Rahul", "Mon · 51 min", "Meeting"},
			{"5", "Standup recap", "Mon · 6 min", "Recap"},
		},
	}
}
