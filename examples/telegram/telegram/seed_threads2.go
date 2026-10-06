package telegram

// seedThreads2 fills the rest of the direct chats' history.
func seedThreads2(m map[string][]Message) {
	m["sofia"] = []Message{
		{ID: "sofia-1", Text: "The slip is ready 🎂", Time: "Mon 14:02"},
		{ID: "sofia-2", Text: "Perfect, I'll pick it up on the way.", Time: "Mon 14:10", Own: true, Read: true},
		{ID: "sofia-3", Text: "Sounds good 👍", Time: "Mon 14:11"},
	}
	m["alex"] = []Message{
		{ID: "alex-1", Text: "Can you send me the file from Friday?", Time: "Sun 17:40"},
		{ID: "alex-2", Text: "Sent the file", Time: "Sun 17:44", Own: true, Read: true},
		{ID: "alex-3", Text: "Got it, thanks!", Time: "Sun 17:51"},
	}
}
