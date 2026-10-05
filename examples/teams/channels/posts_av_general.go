package channels

// avengersPosts builds the Avengers General channel posts.
func avengersPosts() []Post {
	return []Post{
		{
			ID: "av1", Author: "Steve Rogers", Initials: "SR", Color: "blue",
			Time: "9:30 AM", Subject: "Weekly debrief moves to Friday",
			Text:  "Debrief at 4 PM in the war room. Bring field notes and one lesson learned.",
			Likes: 8, Liked: true, Pinned: true,
			Replies: thread("av1", []Reply{
				rep("Natasha Romanoff", "9:35 AM", "I will bring the Zurich extraction notes."),
				rep("Robert Downey Jr.", "9:41 AM", "Slides are ready. Five minutes per mission, tops."),
				rep("Thor", "9:52 AM", "I shall arrive after the Asgard relay check."),
				rep("Bruce Banner", "10:04 AM", "The lab report is in the shared folder already."),
				rep("Clint Barton", "10:12 AM", "Can we keep one thread for open actions?"),
			}),
			Reactions: []Reaction{rx(eHeart, 4), rx(eLaugh, 2)},
		},
		{
			ID: "av2", Author: "Robert Downey Jr.", Initials: "RDJ", Color: "red",
			Time: "8:45 AM", Subject: "Shawarma Friday returns",
			Text:  "Reserved the corner table for Friday at 1 PM. Sign up in the thread.",
			Likes: 6,
			Replies: thread("av2", []Reply{
				rep("Pepper Potts", "8:50 AM", "I put it on the tower calendar."),
				rep("Happy Hogan", "8:58 AM", "Table for twelve confirmed with the kitchen."),
				rep("Wanda Maximoff", "9:10 AM", "Count me in. Extra garlic sauce."),
				rep("Vision", "9:14 AM", "I will attend for the conversation, not the meal."),
			}),
			Reactions: []Reaction{rx(eLove, 5), rx(eSmile, 3), rx(eLaugh, 1)},
		},
	}
}
