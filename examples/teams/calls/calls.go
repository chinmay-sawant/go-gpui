// Package calls is the Calls menu of the Teams example.
package calls

// Data is the calls state the shell prints.
type Data struct {
	Tab        string // "history" | "voicemail"
	Query      string
	Status     string // "" | "calling"
	StatusName string
	Number     string // dialed digits, e.g. "+1 555 0142"
	Active     string
	History    []Call
	Voicemails []Voicemail

	// allHistory and allVoicemails hold every row; the printable
	// lists are the search view of them.
	allHistory    []Call
	allVoicemails []Voicemail
}

// Call is one row in the call history.
type Call struct {
	ID, Name, Initials, Color, Kind, Time, Duration string
	Missed, Incoming                                bool
}

// Voicemail is one voicemail row.
type Voicemail struct {
	ID, Name, Initials, Color, Time, Duration, Transcript string
}

// FromDB rebuilds the calls state from stored rows.
func FromDB(history []Call, voicemails []Voicemail, tab string) Data {
	d := Data{Tab: tab, allHistory: history, allVoicemails: voicemails}
	d.rebuild()

	return d
}

// AllHistory returns every call row, before the search filter.
func (d Data) AllHistory() []Call { return d.allHistory }

// AllVoicemails returns every voicemail row, before the search filter.
func (d Data) AllVoicemails() []Voicemail { return d.allVoicemails }
