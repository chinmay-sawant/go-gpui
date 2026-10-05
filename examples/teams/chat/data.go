package chat

// sampleChats is the read-only chat list every Default clones.
var sampleChats = []Chat{
	{"pepper", "Pepper Potts", "PP", "green", "Pepper Potts: Board review is locked for Thursday.", "9:41 AM", 0, true, false, false},
	{"peter", "Peter Parker", "PP", "red", "You: Firmware v4 is locked. I'll bring the spare cartridges.", "10:04 AM", 1, false, false, false},
	{"avengers", "Avengers Assemble", "AA", "purple", "Steve Rogers: Sokovia relief convoy leaves at dawn.", "9:12 AM", 3, false, false, true},
	{"wakanda", "Wakanda Visit", "WV", "gold", "T'Challa: Bring the shawarma this time.", "8:55 AM", 2, false, false, true},
	{"guardians", "Guardians", "GG", "teal", "Peter Quill: Deal. Also, that is my cassette in your lab.", "Yesterday", 4, false, true, true},
	{"rhodey", "James Rhodes", "JR", "blue", "James Rhodes: Drop test over the base tonight.", "Yesterday", 1, false, false, false},
	{"strange", "Stephen Strange", "SS", "orange", "Stephen Strange: Gate opens at dawn.", "Friday", 0, false, false, false},
	{"wanda", "Wanda Maximoff", "WM", "purple", "Wanda Maximoff: I'll bring the coffee. See you at eight.", "Thursday", 0, false, false, false},
	{"fury", "Nick Fury", "NF", "gray", "Nick Fury: Briefing at 0600 on the deck.", "Wednesday", 0, false, false, false},
}

// Default returns the sample chats with the first chat open.
func Default() Data {
	d := Data{Filter: "all", all: cloneChats(sampleChats)}
	d.rebuild()
	open(&d, d.Chats[0].ID)

	return d
}

// cloneChats copies the sample list so two Defaults never share state.
func cloneChats(src []Chat) []Chat {
	out := make([]Chat, len(src))
	copy(out, src)

	return out
}
