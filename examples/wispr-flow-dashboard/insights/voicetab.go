package insights

// VoiceTabData is the data the insights voice tab prints.
type VoiceTabData struct {
	Name      string
	Status    string
	Line      string
	Checklist []string
}

// DefaultVoiceTab returns the sample insights voice tab content.
func DefaultVoiceTab() VoiceTabData {
	return VoiceTabData{
		Name:   "Subagent Commander",
		Status: "Not set up",
		Line:   "Teach Flow how you sound to improve accuracy.",
		Checklist: []string{
			"More accurate names and technical terms",
			"Consistent punctuation and formatting",
			"The same profile on desktop and mobile",
		},
	}
}
