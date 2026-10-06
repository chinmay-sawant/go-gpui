package telegram

// seedThreads fills the direct chats' message history.
func seedThreads(m map[string][]Message) {
	m["anna"] = seedThreadsAnna()
	m["max"] = []Message{
		{ID: "max-1", Text: "Morning! Did the tests pass overnight?", Time: "09:41"},
		{ID: "max-2", Text: "All green, the flaky one was the clock.", Time: "09:48", Own: true, Read: true},
		{ID: "max-3", Text: "Classic. Thanks for digging in.", Time: "09:52"},
		{ID: "max-4", Text: "The build is green", Time: "10:15", Own: true, Read: true},
	}
	m["mom"] = []Message{
		{ID: "mom-1", Text: "Did you eat something proper today?", Time: "18:20"},
		{ID: "mom-2", Text: "Yes, pasta with the sauce you gave me.", Time: "18:31", Own: true, Read: true},
		{ID: "mom-3", Text: "Good. Call me when you're free", Time: "18:33"},
	}

	seedThreads2(m)
}
