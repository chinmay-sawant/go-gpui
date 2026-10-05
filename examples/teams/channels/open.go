package channels

// openChannel makes one channel active and loads its posts.
func openChannel(d *Data, id string) {
	for i := range d.Teams {
		for j := range d.Teams[i].Channels {
			if d.Teams[i].Channels[j].ID != id {
				continue
			}

			team, ch := d.Teams[i], &d.Teams[i].Channels[j]
			d.ActiveTeam = team.ID
			d.TeamName = team.Name
			d.ActiveChannel = ch.ID
			d.ChannelName = ch.Name
			d.ChannelDesc = describe(ch.ID)
			d.Posts = postsFor(ch.ID)
			d.Tab = "posts"
			ch.Unread = 0

			return
		}
	}
}
