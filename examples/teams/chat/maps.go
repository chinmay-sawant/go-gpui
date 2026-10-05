package chat

// chatPresence maps a direct chat id to its presence dot class.
var chatPresence = map[string]string{
	"pepper":  "online",
	"peter":   "online",
	"rhodey":  "away",
	"strange": "busy",
	"wanda":   "away",
	"fury":    "busy",
}

// chatStatus maps a chat id to the status line in the thread header.
var chatStatus = map[string]string{
	"pepper":    "Available",
	"peter":     "Available",
	"avengers":  "6 members",
	"wakanda":   "3 members",
	"guardians": "4 members",
	"rhodey":    "Away",
	"strange":   "Busy",
	"wanda":     "Away",
	"fury":      "Busy",
}

// Pinned reports whether the open chat is pinned.
func (d Data) Pinned() bool {
	for _, c := range d.Chats {
		if c.ID == d.Active {
			return c.Pinned
		}
	}

	return false
}

// Muted reports whether the open chat is muted.
func (d Data) Muted() bool {
	for _, c := range d.Chats {
		if c.ID == d.Active {
			return c.Muted
		}
	}

	return false
}
