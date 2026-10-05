package channels

// wakandaVibranium builds the Vibranium channel posts.
func wakandaVibranium() []Post {
	return []Post{
		{
			ID: "wkv1", Author: "Shuri", Initials: "SH", Color: "purple",
			Time: "Yesterday", Subject: "Vibranium yield this week",
			Text:  "Six kilograms from the north mine, all above grade. Charts are posted.",
			Likes: 7, Liked: true,
			Replies: thread("wkv1", []Reply{
				rep("T'Challa", "Yesterday", "Route the surplus to the medical wing first."),
				rep("Okoye", "Yesterday", "Mine security found a tunnel. We sealed it."),
				rep("Bruce Banner", "Yesterday", "I would like a sample for the lattice test."),
				rep("Robert Downey Jr.", "8:25 AM", "The suit dampeners need two grams, no more."),
			}),
			Reactions: []Reaction{rx(eWow, 3), rx(eLaugh, 1)},
		},
		{
			ID: "wkv2", Author: "Okoye", Initials: "OK", Color: "orange",
			Time: "Mon", Subject: "Shipment security drill",
			Text:  "Drill runs Thursday at dawn. Two decoy convoys, one real route.",
			Likes: 4,
			Replies: thread("wkv2", []Reply{
				rep("Shuri", "Mon", "I added trackers to all three crates."),
				rep("T'Challa", "Mon", "The river route stays open as the third option."),
				rep("Natasha Romanoff", "Mon", "I can send the stealth checklist we used in Zurich."),
				rep("Clint Barton", "8:55 AM", "Happy to ride shotgun on the decoy run."),
			}),
			Reactions: []Reaction{rx(eHeart, 2), rx(eWow, 1)},
		},
	}
}
