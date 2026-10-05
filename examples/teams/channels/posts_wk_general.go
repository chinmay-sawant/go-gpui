package channels

// wakandaPosts builds the Wakanda R&D General channel posts.
func wakandaPosts() []Post {
	return []Post{
		{
			ID: "wk1", Author: "Shuri", Initials: "SH", Color: "purple",
			Time: "9:05 AM", Subject: "Kimoyo diagnostics update",
			Text:  "The new bead firmware halves sync time. Rolling out to the lab today.",
			Likes: 8, Liked: true, Pinned: true,
			Replies: thread("wk1", []Reply{
				rep("T'Challa", "9:10 AM", "The council will appreciate the faster briefings."),
				rep("Okoye", "9:22 AM", "The Dora Milaje need the update before the border drill."),
				rep("Robert Downey Jr.", "9:35 AM", "Can the beads link to the tower network?"),
				rep("Bruce Banner", "9:48 AM", "I will mirror the protocol in the lab this week."),
			}),
			Reactions: []Reaction{rx(eHeart, 4), rx(eWow, 2)},
		},
		{
			ID: "wk2", Author: "T'Challa", Initials: "TC", Color: "violet",
			Time: "Yesterday", Subject: "Outreach summit in Geneva",
			Text:  "Wakanda speaks at the Geneva summit next month. Draft the panels by Friday.",
			Likes: 6,
			Replies: thread("wk2", []Reply{
				rep("Shuri", "Yesterday", "I will present the water purifier pilot data."),
				rep("Okoye", "Yesterday", "Security plan is ready. Two routes, one decoy."),
				rep("Maria Hill", "8:20 AM", "S.H.I.E.L.D. can share the venue maps."),
				rep("Sam Wilson", "8:44 AM", "Count me in for the relief logistics panel."),
			}),
			Reactions: []Reaction{rx(eLove, 3), rx(eHeart, 2)},
		},
	}
}
