package channels

// shieldPosts builds the S.H.I.E.L.D. Ops General channel posts.
func shieldPosts() []Post {
	return []Post{
		{
			ID: "sh1", Author: "Nick Fury", Initials: "NF", Color: "gray",
			Time: "9:00 AM", Subject: "Helicarrier maintenance window",
			Text:  "Bay three is down for turbine work until Thursday. Route all quinjets to bay two.",
			Likes: 5, Liked: true, Pinned: true,
			Replies: thread("sh1", []Reply{
				rep("Maria Hill", "9:08 AM", "I moved the flight schedule to the board."),
				rep("Carol Danvers", "9:20 AM", "I can cover the orbital watch meanwhile."),
				rep("Robert Downey Jr.", "9:34 AM", "Stark Industries sends the turbine parts tonight."),
				rep("Sam Wilson", "9:47 AM", "Bay two needs a fuel top-up before Friday."),
			}),
			Reactions: []Reaction{rx(eWow, 2), rx(eHeart, 1)},
		},
		{
			ID: "sh2", Author: "Maria Hill", Initials: "MH", Color: "blue",
			Time: "Yesterday", Subject: "Ops roster for November",
			Text:  "The November roster is posted. Swap requests are due by Friday noon.",
			Likes: 3,
			Replies: thread("sh2", []Reply{
				rep("Nick Fury", "Yesterday", "Keep two agents on the night watch at all times."),
				rep("Clint Barton", "Yesterday", "I will take the Thanksgiving week shift."),
				rep("Natasha Romanoff", "Yesterday", "Swap me to the early shift after Tuesday."),
				rep("Carol Danvers", "8:10 AM", "Approved. I signed the updated sheet."),
			}),
			Reactions: []Reaction{rx(eHeart, 3), rx(eSmile, 2)},
		},
	}
}
