package activity

// sampleItems is the read-only feed every Default clones.
var sampleItems = []Item{
	{"a1", "Peter Parker", "PP", "red", "mention", "Avengers Tower › Mission Briefings", "mentioned you in Avengers Tower › Mission Briefings", "Can you review the Mark 85 flight plan before the 2 PM briefing?", "9:42 AM", true, 2, false},
	{"a2", "Maria Hill", "MH", "blue", "reply", "S.H.I.E.L.D. Helicarrier › Command", "replied to your post in S.H.I.E.L.D. Helicarrier › Command", "The orbital traffic numbers hold up, I added the Queens sector split.", "9:18 AM", true, 0, false},
	{"a3", "Shuri", "SH", "purple", "reaction", "Wakanda R&D › Vibranium Allocations", "reacted to your message in Wakanda R&D › Vibranium Allocations", "Love the allocation plan, the sonic dampeners are a smart add.", "8:57 AM", true, 1, false},
	{"a4", "Sam Wilson", "SW", "teal", "follow", "Avengers Compound › Training", "started following Avengers Compound › Training", "You will now see new posts from this channel.", "8:31 AM", false, 0, false},
	{"a5", "Bruce Banner", "BB", "green", "reply", "Stark Industries › Suit Lab", "replied to your comment in Stark Industries › Suit Lab", "The vibranium weave held through the drop test, see the telemetry.", "8:05 AM", false, 3, true},
	{"a6", "Steve Rogers", "SR", "blue", "mention", "Avengers Compound › Training", "mentioned you in Avengers Compound › Training", "The training roster is on slide 12 of the briefing deck.", "7:48 AM", false, 0, false},
	{"a7", "Natasha Romanoff", "NR", "red", "reply", "Avengers Tower › Mission Briefings", "replied to your thread in Avengers Tower › Mission Briefings", "Agreed, let us keep the strike team small for now.", "Yesterday", false, 2, true},
	{"a8", "Thor", "TH", "gold", "reaction", "Asgard › Observatory", "reacted to your file in Asgard › Observatory", "The Bifrost readings are exactly what we needed.", "Yesterday", false, 0, false},
}
