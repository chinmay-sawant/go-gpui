package channels

// guardiansPosts builds the Guardians General channel posts.
func guardiansPosts() []Post {
	return []Post{
		{
			ID: "gd1", Author: "Thor", Initials: "TH", Color: "gold",
			Time: "9:25 AM", Subject: "Milano docked for repairs",
			Text:  "The Guardians' ship is in the Asgard hangar. Rocket found a cracked coolant line.",
			Likes: 6, Liked: true,
			Replies: thread("gd1", []Reply{
				rep("Carol Danvers", "9:30 AM", "I can tow parts from the lunar depot."),
				rep("Peter Parker", "9:44 AM", "Is Groot growing in the cargo bay again?"),
				rep("Thor", "9:58 AM", "Yes. He has claimed the starboard corner."),
				rep("Sam Wilson", "10:10 AM", "Sounds like my old apartment."),
			}),
			Reactions: []Reaction{rx(eLaugh, 5), rx(eHeart, 2)},
		},
		{
			ID: "gd2", Author: "Carol Danvers", Initials: "CD", Color: "gold",
			Time: "Yesterday", Subject: "Refuel window at Knowhere",
			Text:  "The Milano gets a refuel window at Knowhere on Saturday. Two hours only.",
			Likes: 4,
			Replies: thread("gd2", []Reply{
				rep("Thor", "Yesterday", "I shall escort the cargo myself."),
				rep("Nick Fury", "Yesterday", "Keep the manifest logged with S.H.I.E.L.D."),
				rep("Stephen Strange", "Yesterday", "The nearest jump point shifts at dusk."),
				rep("Wanda Maximoff", "8:05 AM", "I can steady the hull if the clamps fail."),
			}),
			Reactions: []Reaction{rx(eWow, 3), rx(eHeart, 2)},
		},
	}
}
