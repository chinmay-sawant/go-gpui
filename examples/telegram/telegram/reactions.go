package telegram

// Reaction is one emoji button in the long-press bar.
type Reaction struct {
	Name  string
	Emoji string
}

// seedReactions returns the bar's emoji. Every one is in the engine's
// curated set, so the buttons paint as color images.
func seedReactions() []Reaction {
	return []Reaction{
		{Name: "like", Emoji: "👍"},
		{Name: "love", Emoji: "❤️"},
		{Name: "laugh", Emoji: "😂"},
		{Name: "wow", Emoji: "😮"},
		{Name: "sad", Emoji: "😢"},
		{Name: "party", Emoji: "🎉"},
	}
}
