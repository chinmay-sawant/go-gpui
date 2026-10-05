package channels

// emojis are the six emoji the bundled font renders in color, in display
// order.
const (
	eHeart = "\u2764\ufe0f"
	eLaugh = "\U0001f602"
	eWow   = "\U0001f62e"
	eSad   = "\U0001f622"
	eLove  = "\U0001f60d"
	eSmile = "\U0001f60a"
)

// emojis is the pickable set in display order.
var emojis = []string{eHeart, eLaugh, eWow, eSad, eLove, eSmile}

// rx builds one starter reaction.
func rx(emoji string, count int) Reaction {
	return Reaction{Emoji: emoji, Count: count}
}
