package telegram

// seedAnnaDinner is the dinner plan Anna's chat list preview quotes.
func seedAnnaDinner() []Message {
	return []Message{
		{ID: "anna-1", Text: "Hey! Are we still on for dinner?", Time: "11:52"},
		{ID: "anna-2", Text: "Yes! I booked the place by the river.", Time: "11:58", Own: true, Read: true},
		{ID: "anna-3", Text: "Amazing 😍 7 pm?", Time: "12:01"},
		{ID: "anna-4", Text: "7 pm works. See you there.", Time: "12:02", Own: true, Read: true},
		{ID: "anna-5", Text: "Perfect, see you at 7", Time: "12:04"},
	}
}
