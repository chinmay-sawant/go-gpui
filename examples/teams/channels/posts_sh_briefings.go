package channels

// shieldBriefings builds the Briefings channel posts.
func shieldBriefings() []Post {
	return []Post{
		{
			ID: "shb1", Author: "Nick Fury", Initials: "NF", Color: "gray",
			Time: "Yesterday", Subject: "Briefing: Loki sighting in Oslo",
			Text:  "Satellite caught a green flare over Oslo. Thor confirms it is not friendly.",
			Likes: 7, Pinned: true,
			Replies: thread("shb1", []Reply{
				rep("Thor", "Yesterday", "It is a decoy. My brother is already off world."),
				rep("Stephen Strange", "Yesterday", "The Sanctum sees no gate activity near Oslo."),
				rep("Wanda Maximoff", "Yesterday", "I felt a spell echo, but it faded fast."),
				rep("Maria Hill", "8:12 AM", "Standing the alert down to level two."),
			}),
			Reactions: []Reaction{rx(eWow, 3), rx(eLaugh, 1)},
		},
		{
			ID: "shb2", Author: "Maria Hill", Initials: "MH", Color: "blue",
			Time: "Mon", Subject: "Clearance updates",
			Text:  "Level seven clearances renew this week. Submit the forms before Thursday.",
			Likes: 4,
			Replies: thread("shb2", []Reply{
				rep("Nick Fury", "Mon", "No exceptions this cycle, even for the Avengers."),
				rep("Carol Danvers", "Mon", "Mine is already in. Space duty needs the access."),
				rep("Clint Barton", "Mon", "Forms done. Can I get the range access back?"),
				rep("Robert Downey Jr.", "8:30 AM", "Sent mine by courier. Paper is safer, apparently."),
			}),
			Reactions: []Reaction{rx(eHeart, 2), rx(eSad, 1)},
		},
	}
}
