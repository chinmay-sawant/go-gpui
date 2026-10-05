package channels

// avengersSuitLab builds the Suit Lab channel posts.
func avengersSuitLab() []Post {
	return []Post{
		{
			ID: "avs1", Author: "Robert Downey Jr.", Initials: "RDJ", Color: "red",
			Time: "Yesterday", Subject: "Mark LXXXV thruster test",
			Text:  "Twelve seconds of hover at 92 percent power. Telemetry is in the shared drive.",
			Likes: 9, Liked: true, Pinned: true,
			Replies: thread("avs1", []Reply{
				rep("Bruce Banner", "Yesterday", "Thermal margins look safe at that burn rate."),
				rep("Peter Parker", "Yesterday", "The stabilizers barely wobbled on the replay."),
				rep("James Rhodes", "Yesterday", "I want that thruster pair on the War Machine frame."),
				rep("Shuri", "8:10 AM", "Try the vibranium dampeners on the ankle joints."),
			}),
			Reactions: []Reaction{rx(eWow, 4), rx(eHeart, 3)},
		},
		{
			ID: "avs2", Author: "Bruce Banner", Initials: "BB", Color: "green",
			Time: "Mon", Subject: "Nano-lattice stress test",
			Text:  "The new lattice held under triple load. Full report is on the lab bench.",
			Likes: 5,
			Replies: thread("avs2", []Reply{
				rep("Robert Downey Jr.", "Mon", "Then we shave two kilos off the chest plate."),
				rep("Shuri", "Mon", "Send me the sample. I will run it against vibranium weave."),
				rep("Peter Parker", "Mon", "Can I watch the next test after school?"),
				rep("Happy Hogan", "8:40 AM", "Only if the lab insurance form is signed first."),
			}),
			Reactions: []Reaction{rx(eHeart, 2), rx(eSmile, 2), rx(eWow, 1)},
		},
	}
}
