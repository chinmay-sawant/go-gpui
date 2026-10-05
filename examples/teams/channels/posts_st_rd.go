package channels

// starkRD builds the Stark Industries R&D channel posts.
func starkRD() []Post {
	return []Post{
		{
			ID: "str1", Author: "Robert Downey Jr.", Initials: "RDJ", Color: "red",
			Time: "Yesterday", Subject: "Repulsor efficiency up 12 percent",
			Text:  "The new nozzle ring gave us a 12 percent gain on the test stand.",
			Likes: 10, Liked: true, Pinned: true,
			Replies: thread("str1", []Reply{
				rep("Bruce Banner", "Yesterday", "The plasma curve stayed clean the whole run."),
				rep("Shuri", "Yesterday", "Vibranium lining would hold that curve longer."),
				rep("Peter Parker", "Yesterday", "I ran the math on the paper. It checks out."),
				rep("James Rhodes", "8:10 AM", "Then the War Machine retrofit moves up a month."),
				rep("Vision", "8:26 AM", "I logged the harmonics for the next build."),
			}),
			Reactions: []Reaction{rx(eWow, 4), rx(eHeart, 3), rx(eLaugh, 1)},
		},
		{
			ID: "str2", Author: "Bruce Banner", Initials: "BB", Color: "green",
			Time: "Mon", Subject: "Arc reactor thermal run",
			Text:  "Ran the reactor at full load for six hours. Coils stayed inside spec.",
			Likes: 7,
			Replies: thread("str2", []Reply{
				rep("Robert Downey Jr.", "Mon", "Good. I want a second run with the older coils."),
				rep("Shuri", "Mon", "Send me the cooling data. I have an idea."),
				rep("Vision", "Mon", "The second run is on the schedule for Friday."),
				rep("Peter Parker", "Mon", "Can I take the readings for the Friday run?"),
			}),
			Reactions: []Reaction{rx(eHeart, 2), rx(eWow, 2)},
		},
	}
}
