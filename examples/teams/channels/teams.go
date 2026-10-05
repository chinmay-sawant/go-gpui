package channels

// defaultTeams returns the five sample teams with fresh slices every call.
func defaultTeams() []Team {
	return []Team{
		{
			ID: "avengers", Name: "Avengers", Initials: "AV", Color: "red", Expanded: true,
			Channels: []Channel{
				{ID: "av-general", Name: "General"},
				{ID: "av-missions", Name: "Missions", Unread: 3},
				{ID: "av-suitlab", Name: "Suit Lab", Unread: 1},
			},
		},
		{
			ID: "wakanda", Name: "Wakanda R&D", Initials: "WK", Color: "purple",
			Channels: []Channel{
				{ID: "wk-general", Name: "General"},
				{ID: "wk-vibranium", Name: "Vibranium", Unread: 2},
				{ID: "wk-outreach", Name: "Outreach"},
			},
		},
		{
			ID: "shield", Name: "S.H.I.E.L.D. Ops", Initials: "SH", Color: "blue",
			Channels: []Channel{
				{ID: "sh-general", Name: "General"},
				{ID: "sh-briefings", Name: "Briefings", Unread: 1},
			},
		},
		{
			ID: "stark", Name: "Stark Industries", Initials: "SI", Color: "gold",
			Channels: []Channel{
				{ID: "st-general", Name: "General"},
				{ID: "st-rd", Name: "R&D", Unread: 4},
			},
		},
		{
			ID: "guardians", Name: "Guardians", Initials: "GD", Color: "teal",
			Channels: []Channel{
				{ID: "gd-general", Name: "General"},
				{ID: "gd-milano", Name: "Milano"},
			},
		},
	}
}
