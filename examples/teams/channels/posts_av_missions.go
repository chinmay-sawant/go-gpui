package channels

// avengersMissions builds the Missions channel posts.
func avengersMissions() []Post {
	return []Post{
		{
			ID: "avm1", Author: "Nick Fury", Initials: "NF", Color: "gray",
			Time: "Yesterday", Subject: "Mission log: Zurich drop",
			Text:  "Extraction window opened at 0300. All agents accounted for, cargo secured.",
			Likes: 7, Liked: true,
			Replies: thread("avm1", []Reply{
				rep("Natasha Romanoff", "Yesterday", "Cover held. The safehouse was already empty."),
				rep("Clint Barton", "Yesterday", "No drones on the roof, so the roof route worked."),
				rep("Sam Wilson", "Yesterday", "Air cover logged two unknown radar sweeps."),
				rep("Steve Rogers", "8:20 AM", "Good run. Debrief notes go to the vault tonight."),
			}),
			Reactions: []Reaction{rx(eWow, 3), rx(eHeart, 2)},
		},
		{
			ID: "avm2", Author: "Natasha Romanoff", Initials: "NR", Color: "red",
			Time: "Mon", Subject: "Next extraction window",
			Text:  "The next window is Thursday 0200 to 0400. Assign roles by Wednesday.",
			Likes: 4,
			Replies: thread("avm2", []Reply{
				rep("Clint Barton", "Mon", "I can take overwatch on the east ridge."),
				rep("Wanda Maximoff", "Mon", "I will hold the perimeter if the lights fail."),
				rep("Thor", "8:05 AM", "Should the Bifrost be a backup exit?"),
				rep("Maria Hill", "8:30 AM", "Yes, but only if the tunnel floods again."),
			}),
			Reactions: []Reaction{rx(eHeart, 3), rx(eSad, 1)},
		},
	}
}
