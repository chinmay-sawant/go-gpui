package channels

// guardiansMilano builds the Milano channel posts.
func guardiansMilano() []Post {
	return []Post{
		{
			ID: "gdm1", Author: "Thor", Initials: "TH", Color: "gold",
			Time: "Yesterday", Subject: "Jump point coordinates",
			Text:  "Charting a new jump point past Titan. The numbers are pinned to the board.",
			Likes: 5, Pinned: true,
			Replies: thread("gdm1", []Reply{
				rep("Carol Danvers", "Yesterday", "Ran the route twice. It is stable enough."),
				rep("Stephen Strange", "Yesterday", "I will anchor the far side with the Sanctum."),
				rep("Peter Parker", "Yesterday", "The math needs one more check on the drift."),
				rep("Nick Fury", "8:15 AM", "Log every jump. We are not losing another ship."),
			}),
			Reactions: []Reaction{rx(eWow, 3), rx(eHeart, 1)},
		},
		{
			ID: "gdm2", Author: "Carol Danvers", Initials: "CD", Color: "gold",
			Time: "Mon", Subject: "Cargo manifest for the relief run",
			Text:  "The Milano carries grain and medicine to the outer colonies this run.",
			Likes: 6,
			Replies: thread("gdm2", []Reply{
				rep("Thor", "Mon", "The hold is full. Even the secret compartment."),
				rep("Sam Wilson", "Mon", "The relief coordinators sent updated drop points."),
				rep("Shuri", "Mon", "Water purifiers are loaded and sealed."),
				rep("Robert Downey Jr.", "8:40 AM", "Fuel cells are on the tower. Do not fly dry."),
			}),
			Reactions: []Reaction{rx(eLove, 3), rx(eHeart, 2)},
		},
	}
}
