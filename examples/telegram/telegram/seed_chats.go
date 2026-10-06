package telegram

// seedChats is the chat list, pinned first, then newest on top.
func seedChats() []Chat {
	return []Chat{
		{
			ID: "anna", Name: "Anna Petrova", Initials: "AP", Color: 0,
			Preview: "Perfect, see you at 7", Time: "12:04",
			Unread: 2, Pinned: true,
		},
		{
			ID: "design", Name: "Design Team", Initials: "DT", Color: 2,
			Preview: "Mira: pushed the new icon set", Time: "11:47",
			Unread: 5, Pinned: true, Group: true,
		},
		{
			ID: "max", Name: "Max Keller", Initials: "MK", Color: 3,
			Preview: "You: the build is green", Time: "10:15",
		},
		{
			ID: "mom", Name: "Mom", Initials: "M", Color: 4,
			Preview: "Call me when you're free", Time: "Yesterday",
			Unread: 1,
		},
		{
			ID: "devs", Name: "Dev Group", Initials: "DG", Color: 5,
			Preview: "Omar: merged #482", Time: "Yesterday",
			Unread: 12, Muted: true, Group: true,
		},
		{
			ID: "sofia", Name: "Sofia Lindqvist", Initials: "SL", Color: 6,
			Preview: "Sounds good 👍", Time: "Mon",
		},
		{
			ID: "alex", Name: "Alex Turner", Initials: "AT", Color: 7,
			Preview: "You: sent the file", Time: "Sun",
		},
	}
}
