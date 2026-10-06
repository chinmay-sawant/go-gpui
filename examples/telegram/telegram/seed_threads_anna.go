package telegram

// seedThreadsAnna fills Anna Petrova's thread. It is long enough to scroll
// several screens on a phone, and it ends with the dinner plan the chat
// list preview quotes.
func seedThreadsAnna() []Message {
	msgs := seedAnnaYesterday()
	msgs = append(msgs, seedAnnaToday()...)

	return append(msgs, seedAnnaDinner()...)
}

// seedAnnaYesterday starts Anna's history the evening before dinner.
func seedAnnaYesterday() []Message {
	return []Message{
		{ID: "anna-a01", Text: "Did you see the pictures from the hike?", Time: "Yesterday 18:02"},
		{ID: "anna-a02", Text: "The lake one is my favourite.", Time: "Yesterday 18:05", Own: true, Read: true},
		{ID: "anna-a03", Text: "I'll send you the full set tonight.", Time: "Yesterday 18:07"},
		{ID: "anna-a04", Text: "No rush, take your time.", Time: "Yesterday 18:09", Own: true, Read: true},
		{ID: "anna-a05", Text: "Also, are you free on Saturday?", Time: "Yesterday 18:15"},
		{ID: "anna-a06", Text: "Saturday works. What's the plan?", Time: "Yesterday 18:18", Own: true, Read: true},
		{ID: "anna-a07", Text: "The market by the station, then lunch?", Time: "Yesterday 18:21"},
		{ID: "anna-a08", Text: "Sold. I haven't been there in ages.", Time: "Yesterday 18:24", Own: true, Read: true},
		{ID: "anna-a09", Text: "They added a bakery stall 😍", Time: "Yesterday 18:30"},
		{ID: "anna-a10", Text: "Now you have my attention.", Time: "Yesterday 18:32", Own: true, Read: true},
		{ID: "anna-a11", Text: "I knew that would work.", Time: "Yesterday 18:35"},
		{ID: "anna-a12", Text: "Guilty. See you Saturday.", Time: "Yesterday 18:41", Own: true, Read: true},
		{ID: "anna-a13", Text: "Night! Talk tomorrow.", Time: "Yesterday 22:14"},
		{ID: "anna-a14", Text: "Night!", Time: "Yesterday 22:15", Own: true, Read: true},
		{ID: "anna-a15", Text: "Bring the umbrella tomorrow, just in case.", Time: "Yesterday 22:17"},
	}
}
