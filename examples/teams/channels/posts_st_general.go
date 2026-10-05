package channels

// starkPosts builds the Stark Industries General channel posts.
func starkPosts() []Post {
	return []Post{
		{
			ID: "st1", Author: "Pepper Potts", Initials: "PP", Color: "pink",
			Time: "9:15 AM", Subject: "Quarterly town hall",
			Text:  "Town hall is Wednesday at 10 AM. Send questions to the thread by Tuesday.",
			Likes: 7, Liked: true, Pinned: true,
			Replies: thread("st1", []Reply{
				rep("Happy Hogan", "9:20 AM", "The auditorium is booked and the coffee is ordered."),
				rep("Robert Downey Jr.", "9:30 AM", "I will demo the repulsor rig at the end."),
				rep("James Rhodes", "9:42 AM", "Keep the demo indoors this time."),
				rep("Maria Hill", "9:55 AM", "Can the town hall stream to the DC office?"),
			}),
			Reactions: []Reaction{rx(eHeart, 4), rx(eSmile, 3)},
		},
		{
			ID: "st2", Author: "Happy Hogan", Initials: "HH", Color: "orange",
			Time: "Yesterday", Subject: "Security badge refresh",
			Text:  "New badges arrive Friday. The old ones stop working at 6 PM sharp.",
			Likes: 4,
			Replies: thread("st2", []Reply{
				rep("Pepper Potts", "Yesterday", "The interns get theirs at the front desk."),
				rep("Robert Downey Jr.", "Yesterday", "Mine should just open everything, thanks."),
				rep("Peter Parker", "Yesterday", "Can the lab interns keep the after-hours badge?"),
				rep("James Rhodes", "8:25 AM", "Only with a signed mentor form. Ask Happy."),
			}),
			Reactions: []Reaction{rx(eLaugh, 2), rx(eHeart, 2)},
		},
	}
}
