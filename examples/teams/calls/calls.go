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

	// allHistory and allVoicemails hold every sample row; the printable
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

// Default returns the sample call history and voicemails. The slices are
// clones, so two Default calls never share mutable state.
func Default() Data {
	d := Data{
		Tab:           "history",
		allHistory:    append([]Call(nil), sampleHistory...),
		allVoicemails: append([]Voicemail(nil), sampleVoicemails...),
	}
	d.rebuild()

	return d
}
