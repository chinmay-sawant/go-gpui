package channels

// wakandaOutreach builds the Outreach channel posts.
func wakandaOutreach() []Post {
	return []Post{
		{
			ID: "wko1", Author: "T'Challa", Initials: "TC", Color: "violet",
			Time: "Yesterday", Subject: "Relief run to Port-au-Prince",
			Text:  "Two flyers leave Friday with medical supplies. Manifest is in the thread.",
			Likes: 9, Liked: true, Pinned: true,
			Replies: thread("wko1", []Reply{
				rep("Shuri", "Yesterday", "Water filters are packed and calibrated."),
				rep("Sam Wilson", "Yesterday", "I can escort the second flyer personally."),
				rep("Robert Downey Jr.", "Yesterday", "The tower covered the fuel cost. Receipts to Pepper."),
				rep("Okoye", "8:15 AM", "Landing clearance is confirmed for both ships."),
			}),
			Reactions: []Reaction{rx(eLove, 4), rx(eHeart, 3)},
		},
		{
			ID: "wko2", Author: "Shuri", Initials: "SH", Color: "purple",
			Time: "Mon", Subject: "School lab kits shipped",
			Text:  "Five hundred lab kits left the workshop this morning. Tracking is live.",
			Likes: 6,
			Replies: thread("wko2", []Reply{
				rep("T'Challa", "Mon", "The teachers will get the lesson plans too."),
				rep("Okoye", "Mon", "Customs paperwork cleared without a hold."),
				rep("Peter Parker", "Mon", "Midtown High would love a kit for the science club."),
				rep("Bruce Banner", "8:05 AM", "Add a radiation badge to each kit. Safety first."),
			}),
			Reactions: []Reaction{rx(eHeart, 3), rx(eSmile, 2)},
		},
	}
}
