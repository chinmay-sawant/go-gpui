package telegram

// seed builds the demo state: the chats, their message threads, and the
// contact list. Everything lives in memory; there is no store.
func seed() *App {
	app := &App{threads: map[string][]Message{}}
	app.chats = seedChats()
	app.contacts = seedContacts()
	seedThreads(app.threads)
	seedGroups(app.threads)
	app.view.Reactions = seedReactions()

	return app
}
