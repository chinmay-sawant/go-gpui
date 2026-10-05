package channels

// FromDB builds the printable state from the stored rows. posts holds every
// channel's posts and becomes the shared source the active slice points at.
func FromDB(teams []Team, posts map[string][]Post, files []FileItem, team, channel, tab string) Data {
	d := Data{
		Teams:         teams,
		allPosts:      posts,
		Files:         files,
		Emojis:        append([]string(nil), emojis...),
		ActiveTeam:    team,
		ActiveChannel: channel,
		Tab:           tab,
		Posts:         posts[channel],
	}

	for i := range teams {
		if teams[i].ID == team {
			d.TeamName = teams[i].Name
		}

		for j := range teams[i].Channels {
			if teams[i].Channels[j].ID == channel {
				d.ChannelName = teams[i].Channels[j].Name
			}
		}
	}

	d.ChannelDesc = describe(channel)

	return d
}
