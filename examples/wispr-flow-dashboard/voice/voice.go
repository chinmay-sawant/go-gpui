// Package voice is the voice profile page: setup status, steps, and details.
package voice

// Data is the data the voice page prints.
type Data struct {
	Steps   []VoiceStep
	Details []VoiceDetail
}

// VoiceStep is one numbered row in the how-it-works panel.
type VoiceStep struct {
	Number string
	Title  string
	Line   string
}

// VoiceDetail is one label and value row in the profile details panel.
// Tag draws the value as a status chip instead of plain text.
type VoiceDetail struct {
	Label string
	Value string
	Tag   bool
}

// Default returns the sample voice-profile page content.
func Default() Data {
	return Data{
		Steps: []VoiceStep{
			{"1", "Record a few samples", "Read short prompts out loud for about a minute."},
			{"2", "Flow learns your voice", "It maps the sounds and names you use most."},
			{"3", "Dictate more accurately", "The profile applies on every device you use."},
		},
		Details: []VoiceDetail{
			{"Name", "Subagent Commander", false},
			{"Language", "English (India)", false},
			{"Status", "Not set up", true},
		},
	}
}
